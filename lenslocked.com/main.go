package main //every go file must have a package name specified

import ( //these are packages imported from the golang standard library
	//formatting strings
	"net/http" //tools for web development
	"text/template"

	"github.com/gorilla/mux"
)

/*
We shouldn't be using a global variable here
but for development purposes we will for now
*/

var homeTemplate *template.Template
var contactTemplate *template.Template

//this will store our template

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func main() {
	//we want to execute the template first before our router
	var err error
	homeTemplate, err = template.ParseFiles("views/home.gohtml")
	if err != nil {
		panic(err) // usually handle errors more gracefully
	}

	contactTemplate, err = template.ParseFiles("views/contact.gohtml")
	if err != nil {
		panic(err) // usually handle errors more gracefully
	}

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)

	http.ListenAndServe(":3000", r) // starts up a local web server using our new gorilla handler
}
