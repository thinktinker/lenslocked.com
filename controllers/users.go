package controllers

import (
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

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}

type Users struct {
	NewView *views.View
}
