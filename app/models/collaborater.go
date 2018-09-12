package models

import (
	"cooptalis-example/app/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

var collaboraterSortFields = []string{"firstname", "lastname", "country", "state", "date_of_entry", "client", "job"}

type Collaborater struct {
	gorm.Model
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	DateOfEntry time.Time `json:"dateOfEntry"`
	Country     string    `json:"country"`
	Job         string    `json:"job"`
	Client      string    `json:"client"`
	State       string    `json:"state"`
}

func InitCollaborater() {
	if len(findAllCollaborater()) > 0 {
		fmt.Println("InitCollaborater : already initialized")
	} else {
		fmt.Println("InitCollaborater : Initialization ...")
		createCollaborater("Stan", "Marsh", time.Now(), "COTE D'IVOIRE", "Java Developer", "ATOS", string(utils.RELOCATION))
		createCollaborater("Kyle", "Broflovski", time.Now(), "COTE D'IVOIRE", "Go Developer", "ATOS", string(utils.IMMIGRATION))
		createCollaborater("Eric", "Cartman", time.Now(), "MAROC", "Android Developer", "SOPRA", string(utils.RELOCATION))
		createCollaborater("Kenny", "Broflovski", time.Now(), "MADAGASCAR", "Python Developer", "SOPRA", string(utils.IMMIGRATION))

	}
}

func createCollaborater(firstname string, lastname string, dateOfEntry time.Time, country string, job string, client string, state string) {
	entry := Collaborater{Firstname: firstname, Lastname: lastname, DateOfEntry: dateOfEntry, Country: country, Job: job, Client: client, State: state}
	Gorm.Create(&entry)
}

func findAllCollaborater() []Collaborater {
	collaboraters := make([]Collaborater, 0)
	err := Gorm.Find(&collaboraters).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return collaboraters

}

func FindCollaboraterById(id uint) (Collaborater, error) {
	var object Collaborater
	err := Gorm.First(&object, id).Error
	return object, err
}

func SearchCollaborater(objects []Collaborater, keyword string, sortField string, sortDirection string) map[string]interface{} {
	if len(sortField) < 1 {
		sortField = collaboraterSortFields[0]
	}
	if len(sortDirection) < 1 {
		sortDirection = "asc"
	}

	var sb strings.Builder
	var result = make(map[string]interface{}, 0)
	sb.WriteString("SELECT * FROM collaboraters")

	if len(strings.Trim(keyword, " ")) > 0 {
		sb.WriteString(" WHERE ")
		for i, v := range collaboraterSortFields {
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
