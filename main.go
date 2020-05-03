package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	fmt.Fprint(w, `<h1> Welcome to my Site </h1>`)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	fmt.Fprint(w, `<h1>To get in touch mail me
		<a href="mailto:diwakarsingh052@gmail.com"> diwakarsingh052@gmail.com </a>
		</h1>`)
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":8080", r)

}
