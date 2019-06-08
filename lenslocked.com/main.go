package main //every go file must have a package name specified

import ( //these are packages imported from the golang standard library
	//formatting strings
	"net/http" //tools for web development

	"github.com/gorilla/mux"
	"github.com/msarah/learngolang/lenslocked.com/views"
)

/*
We shouldn't be using a global variable here
but for development purposes we will for now
*/

var (
	homeView    *views.View
	contactView *views.View
)

//this will store our template

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeView.Template.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := contactView.Template.Execute(w, nil)
	if err != nil {
		panic(err) //panics will clean up later
	}
}

func main() {
	homeView = views.NewView("views/home.gohtml")
	contactView = views.NewView("views/contact.gohtml")
	//we want to execute the template first before our router

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)

	http.ListenAndServe(":3000", r)
	// starts up a local web server using our new gorilla handler
}
