package controllers

import "webdev/views"

func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("Bootstrap", "views/static/home.gohtml"),
		Contact: views.NewView("Bootstrap", "views/static/contact.gohtml"),
	}
}


type Static struct {
	Home    *views.View
	Contact *views.View
}
