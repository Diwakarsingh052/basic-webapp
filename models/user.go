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
	ErrInvalidId  = errors.New("models:Id Provided was invalid")
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
	db := us.db.Where("id=?", id)
	err := first(db, &user)
	return &user, err
}

//looks up the user with email address and returns the user
func (us *UserService) ByEmail(email string) (*User, error) {
	var user User
	db := us.db.Where("email=?", email)
	err := first(db, &user)
	return &user, err

}

//first will query using the provided gorm.DB and it will
//get the first item returned and place into dst
func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error

	if err == gorm.ErrRecordNotFound {
		return ErrorNotFound
	}
	return err
}

//Create will create the user and backfill data
//like id,CreatedAt etc fields
//and return error using gorm when we have one
func (us *UserService) Create(user *User) error {
	return us.db.Create(user).Error
}

//update will update the provided user with all data
func (us *UserService) Update(user *User) error {
	return us.db.Save(user).Error
}

//Delete will delete the data
func (us *UserService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidId
	}
	user := User{Model: gorm.Model{ID: id}}
	return us.db.Delete(&user).Error
}

//Closes UserService Database Connection
func (us *UserService) Close() error {
	return us.db.Close()
}

func (us *UserService) DestructiveReset() error {
	if err := us.db.DropTableIfExists(&User{}).Error; err != nil {
		return err
	}
	return us.AutoMigrate()
}

func (us *UserService) AutoMigrate() error {
	err := us.db.AutoMigrate(&User{}).Error
	if err != nil {
		return err
	}
	return nil
}

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
}
