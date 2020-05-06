package models

import (
	"errors"
	"github.com/jinzhu/gorm"

	"log"
)

type UserService struct {
	db *gorm.DB
}

var (
	ErrorNotFound = errors.New("models:resource not found")
)

func NewUserService(connectionInfo string) (*UserService, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	db.LogMode(true)
	return &UserService{
		db: db,
	}, nil

	return nil, err
}

//By ID will lookup by the id provided
//1- user,nil
//2 -nil,ErrorNotFound
//3 -nil, otherError
func (us *UserService) ByID(id uint) (*User, error) {
	var user User
	err := us.db.Where("id=?", id).First(&user).Error

	switch err {
	case nil:
		return &user, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrorNotFound
	default:
		return nil, err
	}
}

//Create will create the user and backfill data
//like id,CreatedAt etc fields
//and return error using gorm when we have one
func (us *UserService) Create(user *User) error {
	return us.db.Create(user).Error
}

//Closes UserService Database Connection
func (us *UserService) Close() error {
	return us.db.Close()
}

func (us *UserService) DestructiveReset() {
	us.db.DropTableIfExists(&User{})
	us.db.AutoMigrate(&User{})
}

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
}
