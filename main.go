package main

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","text/html")
	//search for *content type* on google for all other content types
	fmt.Fprint(w,"<h1>Welcome to my awesome site!</h1>")
}

func main() {

	http.HandleFunc("/",handleFunc)
	http.ListenAndServe(":8080",nil)

}




