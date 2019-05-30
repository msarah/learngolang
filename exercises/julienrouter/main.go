package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my site!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "Hello, %s!\n", ps.ByName("name"))
	fmt.Fprint(w, "To get in touch, please send an email "+
		"to <a href=\"mailto:support@lenslocked.com\">"+
		"support@lenslocked.com</a>.")
}

func main() {
	r := httprouter.New()
	r.GET("/", home)
	r.GET("/:name", contact)

	log.Fatal(http.ListenAndServe(":3000", r))
}
