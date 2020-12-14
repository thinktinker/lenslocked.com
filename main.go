package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"lenslocked.com/controllers"
)

func main() {

	//1. Create an instance of a new Users controller
	//that returns the address of a struct &Users instance
	//also, added a new controller for static pages
	userC := controllers.NewUsers()
	staticC := controllers.NewStatic()

	// Note: gorilla mux's Handle interface takes
	// in a default ServeHTTP method as a second argument.
	// As long as ServeHTTP is applied in view.go, the below would work
	r := mux.NewRouter()
	r.Handle("/", staticC.HomeView).Methods("GET")
	r.Handle("/contact", staticC.ContactView).Methods("GET")
	r.Handle("/faq", staticC.FaqView).Methods("GET")

	// 1.1 Use the user controller to handle routes for a GET request
	r.HandleFunc("/signup", userC.New).Methods("GET")

	// 2 Set the /signup to lookout for a POST request
	r.HandleFunc("/signup", userC.Create).Methods("POST")

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
