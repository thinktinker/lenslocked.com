package controllers

import "lenslocked.com/views"

func NewStatic() *Static {
	return &Static{
		HomeView:    views.NewView("bootstrap", "static/home"),
		ContactView: views.NewView("bootstrap", "static/contact"),
		FaqView:     views.NewView("bootstrap", "static/faq"),
	}
}

type Static struct {
	HomeView     *views.View
	ContactView  *views.View
	FaqView      *views.View
	NotFoundView *views.View
}
