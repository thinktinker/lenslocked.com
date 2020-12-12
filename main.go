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
	catchErr(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html") //or you can you text/plain
	catchErr(contactView.Render(w, nil))
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	catchErr(faqView.Render(w, nil))
}

func notfound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	catchErr(notfoundView.Render(w, nil))

}

func catchErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	faqView = views.NewView("bootstrap", "views/faq.gohtml")
	notfoundView = views.NewView("bootstrap", "views/notfound.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.NotFoundHandler = http.HandlerFunc(notfound)
	http.ListenAndServe(":3000", r)
}
