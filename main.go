package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"lenslocked.com/views"
)

var (
	homeView     *views.View
	contactView  *views.View
	faqView      *views.View
	notfoundView *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeView.Template.Execute(w, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html") //or you can you text/plain

	//you need to call Template.Execute to access the struct's View Template
	//the second argurment is set as nil as there's no data to pass to it
	if err := contactView.Template.Execute(w, nil); err != nil {
	}
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	//you need to call Template.Execute to access the struct's View Template
	//the second argurment is set as nil as there's no data to pass to it
	if err := faqView.Template.Execute(w, nil); err != nil {
		panic(err)
	}
}

func notfound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)

	//you need to call Template.Execute to access the struct's View Template
	//the second argurment is set as nil as there's no data to pass to it
	if err := notfoundView.Template.Execute(w, nil); err != nil {
		panic(err)
	}
}

func main() {

	homeView = views.NewView("views/home.gohtml")
	contactView = views.NewView("views/contact.gohtml")
	faqView = views.NewView("views/faq.gohtml")
	notfoundView = views.NewView("views/notfound.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.NotFoundHandler = http.HandlerFunc(notfound)
	http.ListenAndServe(":3000", r)
}
