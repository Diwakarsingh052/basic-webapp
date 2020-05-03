package main

import (
	"html/template"
	"log"
	"os"
)

type user struct {
	Name string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml") //paths are relative it means that this code will look for this file in this same folder
	if err != nil {
		log.Println(err)
		return
	}

	data := user{
		Name: "Diwakar",
	}

	err = t.Execute(os.Stdout, data) //stdout prints to the terminal
	if err != nil {
		panic(err)
	}
	data.Name = "Dev"
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

//more example Cast_4 Ex3