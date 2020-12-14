package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"lenslocked.com/controllers"
	"lenslocked.com/views"
)

var (
	notfoundView *views.View
)

// Handler functions (or Actions) START here
// ****************************************

func notfound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	must(notfoundView.Render(w, nil))
}

// ****************************************
// Handler functions (or Actions) END here

func main() {

	//1. Create an instance of a new Users controller
	//   that returns the address of a struct &Users instance
	//   UPDATE: All static pages are now served through a a controller
	userC := controllers.NewUsers()
	staticC := controllers.NewStatics()

	notfoundView = views.NewView("bootstrap404", "views/static/notfound.gohtml")

	r := mux.NewRouter()

	// UPDATE: Use Gorilla's Handle interface to route static pages
	// The Handle interface has a default function ServeHTTP
	// Therefore was implemented in view.go as well
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.Handle("/faq", staticC.Faq).Methods("GET")

	// 1.1 Use the user controller to handle routes for a GET request
	r.HandleFunc("/signup", userC.New).Methods("GET")

	// 2 Set the /signup to lookout for a POST request
	r.HandleFunc("/signup", userC.Create).Methods("POST")

	r.NotFoundHandler = http.HandlerFunc(notfound)
	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

//How the template rendering improved

//Opton 1:
//you need to call Template.Execute to access the struct's View Template
//the second argurment is set as nil as there's no data to pass to it
// if err := homeView.Template.Execute(w, nil); err != nil {
// 	panic(err)
// }

//Option 2:
//Improved: use ExecuteTemplate to render home.gohtml's "yield" template set in bootstrap.gohtml
// if err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil); err != nil {
// 	panic(err)
// }

//Option 3:
// Improved again: move the rendering process ExecuteTemplate to views.go
// The must function is used to render the errors in the console
// must(homeView.Render(w, nil))
