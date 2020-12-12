package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"lenslocked.com/controllers"
	"lenslocked.com/views"
)

var (
	homeView     *views.View
	contactView  *views.View
	faqView      *views.View
	notfoundView *views.View
)

// Handler functions (or Actions) START here
// ****************************************

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html") //or you can you text/plain
	must(contactView.Render(w, nil))
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(faqView.Render(w, nil))
}

func notfound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	must(notfoundView.Render(w, nil))
}

// ****************************************
// Handler functions (or Actions) END here

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	faqView = views.NewView("bootstrap", "views/faq.gohtml")

	//1. Create an instance of a new Users controller
	//that returns the address of a struct &Users instance
	userC := controllers.NewUsers()

	notfoundView = views.NewView("bootstrap404", "views/notfound.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)

	// 2. use the user controller to handle the routing
	r.HandleFunc("/signup", userC.New)

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
