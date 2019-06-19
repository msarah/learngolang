package controllers

import (
	"net/http"

	"github.com/gorilla/schema"
)

//parseForm receives the necessary form via the
//http.Request object and takes the "dst"
//argument which declares where the data from the
//form should go, which can be of any type. This means
//we have to scan for errors because a type that cannot
//be decoded could be passed in
func parseForm(r *http.Request, dst interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	//create decoder with schema package
	dec := schema.NewDecoder()

	//use the decoder to take the data from the request form,
	//convert to the correct data types and insert into our
	//new SignupForm struct
	if err := dec.Decode(dst, r.PostForm); err != nil {
		return err
	}
	return nil
}
