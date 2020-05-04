package controllers

import (
	"github.com/gorilla/schema"
	"net/http"
)

//Parse the form and decode it in struct.
//Created by me
func parseForm(r *http.Request, dst interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	dec := schema.NewDecoder()

	if err := dec.Decode(dst, r.PostForm); err != nil {
		return err
	}
	return nil

}
