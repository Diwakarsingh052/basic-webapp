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
		NewView:   views.NewView("bootstrap", "users/new"),
		LoginView: views.NewView("bootstrap", "users/login"),
		us:        us,
	}
}

type Users struct {
	NewView   *views.View
	LoginView *views.View
	us        *models.UserService
}

//New is used to render the form where user can create a new user account
// GET /signup
//func (u *Users) New(w http.ResponseWriter, r *http.Request) {
//	if err := u.NewView.Render(w, nil); err != nil {
//		log.Println(err)
//	}
//}

// Create is used to process the sign up form
//when user submits it.

//Post /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		log.Println(err)
	}
	user := models.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}
	err := u.us.Create(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, user)

}

type LoginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	form := LoginForm{}
	err := parseForm(r, &form)
	if err != nil {
		panic(err)
	}

	user, err := u.us.Authenticate(form.Email, form.Password)
	if err != nil {
		switch err {
		case models.ErrorNotFound:
			fmt.Fprintln(w, "Invalid Email Address")
		case models.ErrInvalidIdPassword:
			fmt.Fprintln(w, "Invalid Password")
		default:
			fmt.Fprintln(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	cookie := http.Cookie{
		Name:  "email", //key
		Value: user.Email,
	}

	http.SetCookie(w, &cookie)

	fmt.Fprintln(w, user)
}

func (u *Users) CookieTest(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("email")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Email:", cookie.Name)
	fmt.Fprintln(w, cookie)
}
