package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"webdev/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "diwakar"
	password = "root"
	dbname   = "website"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()

	//us.DestructiveReset()
	us.AutoMigrate()
	user := models.User{
		Model:    gorm.Model{},
		Name:     "Diwakar",
		Email:    "diwakar@gmail.com",
		Password: "jon",
		Remember: "abc123",
	}

	us.Create(&user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", user)

	user2, err := us.ByRemember("abc123")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", user2)

}
