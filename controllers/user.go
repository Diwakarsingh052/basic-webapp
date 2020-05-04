package controllers

import (
	"fmt"
	"net/http"
	"webdev/views"
)

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

type Users struct {
	NewView *views.View
}

//New is used to render the form where user can create a new user account
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

// Create is used to process the sign up form
//when user submits it.

//Post /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is a sample response ")
}
