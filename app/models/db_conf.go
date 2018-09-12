package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

var Gorm *gorm.DB

func InitDb() {
	username := revel.Config.StringDefault("db.user", "")
	password := revel.Config.StringDefault("db.pass", "")
	dbName := revel.Config.StringDefault("db.name", "")
	dbHost := revel.Config.StringDefault("db.host", "")
	InitDB(username, password, dbName, dbHost)
}

func InitDB(username string, password string, dbName string, dbHost string) {

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	Gorm = conn
	Gorm.Debug().AutoMigrate(Client{}, Collaborater{}, Menu{}, User{})
}
