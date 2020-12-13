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
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
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

	// 1. r.ParseForm is required to put the results of
	//    a POST, PUT or PATCH method into r.PostForm or r.Form
	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	// 2. Postform's results are derived as key value pairs
	//    r.PostForm = map[string][]string
	fmt.Fprintln(w, r.PostForm["email"])
	fmt.Fprintln(w, r.PostForm["password"])

	// 3. PostFormValue does not need r.ParseForm
	//    It ignores errors and returns the first value of a named component
	// fmt.Fprintln(w, r.PostFormValue("email"))
	// fmt.Fprintln(w, r.PostFormValue("password"))

}

type Users struct {
	NewView *views.View
}
