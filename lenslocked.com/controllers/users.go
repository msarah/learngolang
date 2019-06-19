package controllers

import (
	"fmt"
	"net/http"

	"github.com/msarah/learngolang/lenslocked.com/views"
)

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
	}
}

//this is our user controller
type Users struct {
	NewView *views.View //this is here so we can easily call the render method
}

type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

//New is a method function which means it has access to
//the Users controller and actions which allows us
//to render the view necessary for a new user to sign up
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
