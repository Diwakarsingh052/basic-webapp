package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"webdev/controllers"
)

func notfound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Oops Page not Found</h1>")
}

func main() {

	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers()
	r := mux.NewRouter()

	r.NotFoundHandler = http.HandlerFunc(notfound)
	r.Handle("/", staticC.Home).Methods("GET") //METHOD- here we specify which method we are going to use for a particular endpoint. By default all
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	http.ListenAndServe(":8080", r)

}
