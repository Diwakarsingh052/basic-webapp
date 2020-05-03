package main

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	//search for *content type* on google for all other content types

	if r.URL.Path == "/" {
		fmt.Fprint(w, `<h1> Welcome to my Site </h1>`)
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, `<h1>To get in touch mail me
		<a href="mailto:diwakarsingh052@gmail.com"> diwakarsingh052@gmail.com </a>
		</h1>`)
	}

}

func main() {

	http.HandleFunc("/", handleFunc)
	http.ListenAndServe(":8080", nil)

}
