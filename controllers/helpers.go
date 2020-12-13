package controllers

import (
	"net/http"

	"github.com/gorilla/schema"
)

func parseForm(r *http.Request, dst interface{}) error {

	// 1. r.ParseForm is needed; it put results of
	//    a POST, PUT or PATCH method into r.PostForm or r.Form
	if err := r.ParseForm(); err != nil {
		return err
	}

	// 2. Instead of using PostForm to call on individual form values
	//    e.g fmt.Fprintln(w, r.PostForm['email'])
	//    gorilla schema's Decoder is used to grab all the POST values
	dec := schema.NewDecoder()

	// 3. dst or destination is sent as a pointer to this function
	//    and values decoded from r.PostForm are stored in the
	//    actual address and not just a copy
	if err := dec.Decode(dst, r.PostForm); err != nil {
		return err
	}

	return nil
}
