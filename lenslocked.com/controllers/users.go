package controllers

import (
	"fmt"
	"net/http"

	"github.com/msarah/learngolang/lenslocked.com/views"
)

//this returns a new users controller for use
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
	}
}

//this is our user controller
type Users struct {
	NewView *views.View //this is here so we can easily call the render method
}

//this form is where we hold the data that we receive back from the user
//when they sign up
type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

//New is a method function which means it has access to
//the Users controller and actions which allows us
//to render the view necessary for a new user to sign up
//this implements the handler interface which is what we pass
//to our router to render the view
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, "Email is:", form.Email)
	fmt.Fprintln(w, "Password is:", form.Password)
}
