package controllers

import (
	"fmt"
	"log"
	"net/http"
	"webdev/models"
	"webdev/views"
)

type SignupForm struct {
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func NewUsers(us *models.UserService) *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
		us:      us,
	}
}

type Users struct {
	NewView *views.View
	us      *models.UserService
}

//New is used to render the form where user can create a new user account
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		log.Println(err)
	}
}

// Create is used to process the sign up form
//when user submits it.

//Post /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		log.Println(err)
	}
	user := models.User{
		Name:  form.Name,
		Email: form.Email,
	}
	err := u.us.Create(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, form)

}
