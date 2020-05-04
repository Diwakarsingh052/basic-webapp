package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"webdev/controllers"
	"webdev/views"
)

var (
	homeView    *views.View
	contactView *views.View

)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	if err := homeView.Render(w, nil); err != nil {
		//nil because we are not writing any data
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	if err := contactView.Render(w, nil); err != nil {
		panic(err)
	}
}

func notfound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Oops Page not Found</h1>")
}


func main() {

	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")

	usersC := controllers.NewUsers()
	r := mux.NewRouter()

	r.NotFoundHandler = http.HandlerFunc(notfound)
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/signup", usersC.New)
	http.ListenAndServe(":8080", r)

}
