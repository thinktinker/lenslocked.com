package controllers

import "lenslocked.com/views"

func NewStatics() *Statics {
	return &Statics{
		Home:    views.NewView("bootstrap", "views/static/home.gohtml"),
		Contact: views.NewView("bootstrap", "views/static/contact.gohtml"),
		Faq:     views.NewView("bootstrap", "views/static/faq.gohtml"),
	}
}

type Statics struct {
	Home    *views.View
	Contact *views.View
	Faq     *views.View
}
