package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"webdev/hash"
	"webdev/rand"
)

type UserService struct {
	UserDB
}

type userValidator struct {
	UserDB
}

type userGorm struct {
	db   *gorm.DB
	hmac hash.HMAC
}

var _UserDB = &userGorm{}
var (
	ErrorNotFound        = errors.New("models:resource not found")
	ErrInvalidId         = errors.New("models:Id Provided was invalid")
	ErrInvalidIdPassword = errors.New("models:Password Provided was invalid")
)

const userPwPepper = "secret-random-string"
const hmacSecretKey = "secret-hmac-key"

//UserDB is used to interact with users database
type UserDB interface {
	//Methods for querying for single users
	ByID(id uint) (*User, error)
	ByEmail(email string) (*User, error)
	ByRemember(token string) (*User, error)

	//Methods for altering users
	Create(user *User) error
	Update(user *User) error
	Delete(id uint) error

	//Close Connection
	Close() error

	//Migration helpers
	AutoMigrate() error
	DestructiveReset() error
}

func newUserGorm(connectionInfo string) (*userGorm, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	db.LogMode(true)
	hmac := hash.NewHMAC(hmacSecretKey)
	return &userGorm{
		db:   db,
		hmac: hmac,
	}, nil

	return nil, err
}

func NewUserService(connectionInfo string) (*UserService, error) {
	ug, err := newUserGorm(connectionInfo)
	if err != nil {
		return nil, err
	}

	return &UserService{
		UserDB: &userValidator{UserDB: ug},
	}, nil

}

//By ID will lookup by the id provided
//1- user,nil
//2 -nil,ErrorNotFound
//3 -nil, otherError
func (ug *userGorm) ByID(id uint) (*User, error) {
	var user User
	db := ug.db.Where("id=?", id)
	err := first(db, &user)
	return &user, err
}

//looks up the user with email address and returns the user
func (ug *userGorm) ByEmail(email string) (*User, error) {
	var user User
	db := ug.db.Where("email=?", email)
	err := first(db, &user)
	return &user, err

}

//By remember looks up a user with the given remember token
//and return that user.This method will handle hashing the token
//for us.
func (ug *userGorm) ByRemember(token string) (*User, error) {
	var user User
	rememberHash := ug.hmac.Hash(token)
	err := first(ug.db.Where("remember_hash=?", rememberHash), &user)
	if err != nil {
		return nil, err
	}
	return &user, err
}

// Authenticate can be used to authenticate a user
func (us *UserService) Authenticate(email, password string) (*User, error) {
	foundUser, err := us.ByEmail(email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.PasswordHash), []byte(password+userPwPepper))

	if err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			return nil, ErrInvalidIdPassword
		default:
			return nil, err
		}
	}

	return foundUser, nil

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
func (ug *userGorm) Create(user *User) error {
	pwBytes := []byte(user.Password + userPwPepper)
	hashedByte, err := bcrypt.GenerateFromPassword(pwBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedByte)
	user.Password = ""
	if user.Remember == "" {
		token, err := rand.RememberToken()
		if err != nil {
			return err
		}
		user.Remember = token
	}
	user.RememberHash = ug.hmac.Hash(user.Remember)

	return ug.db.Create(user).Error
}

//update will update the provided user with all data
func (ug *userGorm) Update(user *User) error {
	if user.Remember != "" {
		user.RememberHash = ug.hmac.Hash(user.Remember)
	}
	return ug.db.Save(user).Error
}

//Delete will delete the data
func (ug *userGorm) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidId
	}
	user := User{Model: gorm.Model{ID: id}}
	return ug.db.Delete(&user).Error
}

//Closes userGorm Database Connection
func (ug *userGorm) Close() error {
	return ug.db.Close()
}

func (ug *userGorm) DestructiveReset() error {
	if err := ug.db.DropTableIfExists(&User{}).Error; err != nil {
		return err
	}
	return ug.AutoMigrate()
}

func (ug *userGorm) AutoMigrate() error {
	err := ug.db.AutoMigrate(&User{}).Error
	if err != nil {
		return err
	}
	return nil
}

type User struct {
	gorm.Model
	Name         string
	Email        string `gorm:"not null;unique_index"`
	Password     string `gorm:"-"`
	PasswordHash string `gorm:"not null"`
	Remember     string `gorm:"-"`
	RememberHash string `gorm:"not null;unique_index"`
}
