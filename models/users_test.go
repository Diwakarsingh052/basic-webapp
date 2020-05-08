package models

import (
	"fmt"
	_ "github.com/lib/pq"
	"testing"
)

func testingUserService() (*userGorm, error) {
	const (
		host     = "localhost"
		port     = 5432
		user     = "diwakar"
		password = "root"
		dbname   = "website_test"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	us, err := NewUserService(psqlInfo)
	if err != nil {
		return nil, err
	}

	us.db.LogMode(false)
	us.DestructiveReset()
	//clear the users table between test
	return us, nil
}

func TestUserService(t *testing.T) {
	us, err := testingUserService()
	if err != nil {
		t.Fatal(err)
	}
	user := User{
		Name:  "Diwakar",
		Email: "diwakar@gmail.com",
	}

	err = us.Create(&user)
	if err != nil {
		t.Fatal(err)
	}

	if user.ID == 0 {
		t.Errorf("Expected Id > 0 Received %d", user.ID)
	}
}
