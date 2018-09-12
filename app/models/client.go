package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

var clientSortFields = []string{"title", "email"}

type Client struct {
	gorm.Model
	Title string
	Email string
}

func InitClient() {
	if len(FindAllClient()) > 0 {
		fmt.Println("InitClient : already initialized")
	} else {
		fmt.Println("InitClient : Initialization ...")
		createClient("ATOS", "client@atos.fr")
		createClient("SOPRA", "client@sopra.fr")

	}

}

func createClient(title string, email string) Client {
	entry := Client{Title: title, Email: email}
	Gorm.Create(&entry)
	return entry
}

func FindAllClient() []Client {
	clients := make([]Client, 0)
	err := Gorm.Find(&clients).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return clients

}

func FindClientById(id uint) (Client, error) {
	var object Client
	err := Gorm.First(&object, id).Error
	return object, err
}

func SearchClient(objects []Client, keyword string, sortField string, sortDirection string) map[string]interface{} {
	if len(sortField) < 1 {
		sortField = clientSortFields[0]
	}
	if len(sortDirection) < 1 {
		sortDirection = "asc"
	}

	var sb strings.Builder
	var result = make(map[string]interface{}, 0)
	sb.WriteString("SELECT * FROM clients")

	if len(strings.Trim(keyword, " ")) > 0 {
		sb.WriteString(" WHERE ")
		for i, v := range clientSortFields {
			if i != 0 {
				sb.WriteString(" OR ")
			}
			sb.WriteString(v)
			sb.WriteString(" LIKE ")
			sb.WriteString("'%")
			sb.WriteString(keyword)
			sb.WriteString("%'")
		}
	}
	sb.WriteString(" ORDER BY ")
	sb.WriteString(sortField)
	sb.WriteString(" ")
	sb.WriteString(sortDirection)
	Gorm.Raw(sb.String()).Scan(&objects)

	result["objects"] = objects
	result["count"] = len(objects)

	return result
}
