package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var Gorm *gorm.DB

func InitDb() {
	conn, err := gorm.Open("sqlite3", "cooptalis.db")
	if err != nil {
		fmt.Print(err)
	}

	Gorm = conn
	Gorm.Debug().AutoMigrate(Client{}, Collaborater{}, Menu{}, User{})
}

//
