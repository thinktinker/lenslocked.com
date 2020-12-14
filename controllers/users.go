package controllers

import (
	"fmt"
	"net/http"

	"lenslocked.com/views"
)

// NewUsers is used to create a new Users controller
// THis function will panic if the templates are not
// parsed correctly, and should only be used during the
// initial setup.

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
	}
}

// New is used to render a form for a user to create an account
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

// Create is used to process a signup form when a user submits it.
// This is used to create a new user account.
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {

	// Created a helper function based on DRY principles
	// to parse form data, the form should be sent as a pointer
	// so that it may be referenced here
	var form SignUpForm

	if err := parseForm(r, &form); err != nil {
		panic(err)
	}

	fmt.Fprintln(w, "The results captured is:")
	fmt.Fprintln(w, form)
}

type Users struct {
	NewView *views.View
}

type SignUpForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}
