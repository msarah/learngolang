package controllers

import "github.com/msarah/learngolang/lenslocked.com/views"

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

type Users struct {
	NewView *views.View //this is here so we can easily call the render method
}
