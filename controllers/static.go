package controllers

import "lenslocked.com/views"

func NewStatic() *Static {
	return &Static{
		HomeView:    views.NewView("bootstrap", "views/static/home.gohtml"),
		ContactView: views.NewView("bootstrap", "views/static/contact.gohtml"),
		FaqView:     views.NewView("bootstrap", "views/static/faq.gohtml"),
	}
}

type Static struct {
	HomeView     *views.View
	ContactView  *views.View
	FaqView      *views.View
	NotFoundView *views.View
}
