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

	//Opton 1:
	//you need to call Template.Execute to access the struct's View Template
	//the second argurment is set as nil as there's no data to pass to it
	// if err := homeView.Template.Execute(w, nil); err != nil {
	// 	panic(err)
	// }

	//Option 2:
	//Improved: use ExecuteTemplate to render home.gohtml's "yield" template set in bootstrap.gohtml
	if err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html") //or you can you text/plain

	//Option 1:
	//you need to call Template.Execute to access the struct's View Template
	//the second argurment is set as nil as there's no data to pass to it
	// if err := contactView.Template.Execute(w, nil); err != nil {
	// 	panic(err)
	// }

	//Option 2:
	//Improved: use ExecuteTemplate to render contact.gohtml's "yield" template set in boostrap.gohtml
	if err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil); err != nil {
		panic(err)
	}
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	//Option 1:
	//you need to call Template.Execute to access the struct's View Template
	//the second argurment is set as nil as there's no data to pass to it
	// if err := faqView.Template.Execute(w, nil); err != nil {
	// 	panic(err)
	// }

	//Option 2:
	//Improved: Use ExecuteTemplate to render faq.gohtml's "yield" template set in bootstrap.gohtml
	if err := faqView.Template.ExecuteTemplate(w, faqView.Layout, nil); err != nil {
		panic(err)
	}
}

func notfound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)

	//Option 1
	//you need to call Template.Execute to access the struct's View Template
	//the second argurment is set as nil as there's no data to pass to it
	// if err := notfoundView.Template.Execute(w, nil); err != nil {
	// 	panic(err)
	// }

	//Option 2
	if err := notfoundView.Template.ExecuteTemplate(w, notfoundView.Layout, nil); err != nil {
		panic(err)
	}
}

func main() {

	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	faqView = views.NewView("bootstrap", "views/faq.gohtml")
	notfoundView = views.NewView("bootstrap404", "views/notfound.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.NotFoundHandler = http.HandlerFunc(notfound)
	http.ListenAndServe(":3000", r)
}
