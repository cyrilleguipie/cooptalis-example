package models

import (
	"cooptalis-example/app/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

type Menu struct {
	gorm.Model
	Title    string `json:"title"`
	Type     string `json:"type"`
	ApiURL   string `json:"apiURL"`
	IconURL  string `json:"iconURL"`
	ParentID uint   `json:"parentID"`
}

var menuSortFields = []string{"title", "type"}

func InitMenu() {

	if len(findAllMenu()) > 0 {
		fmt.Println("InitMenu : already initialized")
	} else {

		fmt.Println("InitMenu : Initialization ...")
		// url = /foo?sort=asc&active=1
		//Parent menu
		MenuCollaborater := createMenu("Collaborateurs", string(utils.LIST), "/collaborateurs", "http://i.imgur.com/DRuOeD0.png", 0)
		createMenu("Relocation", string(utils.LIST), "/collaborateurs?k=relocation", "http://i.imgur.com/sRK4tw3.png", MenuCollaborater.ID)
		createMenu("Immigration", string(utils.LIST), "/collaborateurs?k=immigration", "http://i.imgur.com/EDCW4OD.png", MenuCollaborater.ID)

		//Children Menu relocation https://imgur.com/sRK4tw3 immigration https://imgur.com/EDCW4OD
		createMenu("Clients", string(utils.LIST), "/clients", "http://i.imgur.com/BuaklRB.png", 0)

	}

}

func createMenu(title string, _type string, apiURL string, iconURL string, parentId uint) Menu {
	entry := Menu{Title: title, Type: _type, ApiURL: apiURL, IconURL: iconURL, ParentID: parentId}
	Gorm.Create(&entry)
	return entry
}

func findAllMenu() []Menu {
	menus := make([]Menu, 0)
	err := Gorm.Find(&menus).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return menus

}

func SearchMenu(objects []Menu, keyword string, sortField string, sortDirection string) map[string]interface{} {
	if len(sortField) < 1 {
		sortField = menuSortFields[0]
	}
	if len(sortDirection) < 1 {
		sortDirection = "asc"
	}

	var sb strings.Builder
	var result = make(map[string]interface{}, 0)
	sb.WriteString("SELECT * FROM menus")

	if len(strings.Trim(keyword, " ")) > 0 {
		sb.WriteString(" WHERE ")
		for i, v := range menuSortFields {
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
