package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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
	//us.DestructiveReset()
	//user := models.User{
	//	Name:  "Dev",
	//	Email: "dev@g.com",
	//}
	//
	//err = us.Create(&user)

	user, err := us.ByID(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
	defer us.Close()

}
