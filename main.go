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



func notfound(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w,"<h1>Oops Page not Found</h1>")
}

func main() {

	r := mux.NewRouter()
	// not Found handler needs http.Handler which is implemented by
	//handlerfunc type  and handlerfunc type implements w http.ResponseWriter, r *http.Request
	//check docs for more info // video cast 3.2_EX_2
	r.NotFoundHandler= http.HandlerFunc(notfound)
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":8080", r)

}
