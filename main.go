package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html") //or you can you text/plain
	// fmt.Fprint(w, r.URL.Path)                //r.URL.Path shows the current URL Path requested

	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my web site</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "Get in touch with us. Send an email to <a href=\"mailto:support@lenslocked.com\">lenslocked.com</a>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>We could not find the page you were looking for.</h1>")
	}

}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
