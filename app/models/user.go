package models

import (
	"cooptalis-example/app/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Email        string `json:"email"`
	PasswordHash []byte
	Role         string `json:"role"`
	Token        string
}

var userSortFields = []string{"email", "role"}

func InitUser() {
	if len(FindAllUser()) > 0 {
		fmt.Println("InitUser : already initialized")
	} else {
		fmt.Println("InitUser : Initialization ...")
		createUser("admin@cooptalis.com", "admin", string(utils.ADMIN), "ADMIN123")
		createUser("member@cooptalis.com", "member", string(utils.MEMBER), "MEMBER123")
		createUser("anonymous@cooptalis.com", "anonymous", string(utils.ANONYMOUS), "AMONYMOUS123")

	}

}

func createUser(email string, password string, role string, token string) User {
	entry := User{Email: email, Role: role, Token: token}
	entry.SetPassword(password)
	Gorm.Create(&entry)
	return entry
}

func (user *User) SetPassword(password string) {
	BcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user.PasswordHash = BcryptPassword
}

func FindByToken(token string) (*User, error) {
	users := make([]User, 0)
	err := Gorm.Where("token = ?", token).Find(&users).Error
	if err != nil || len(users) < 1 {
		fmt.Println(err)
		return nil, err
	} else {
		return &users[0], nil
	}

}

func FindAllUser() []User {
	users := make([]User, 0)
	err := Gorm.Find(&users).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return users

}

//TODO: JWT, Validation(Email, Password), Register, Auth ...
