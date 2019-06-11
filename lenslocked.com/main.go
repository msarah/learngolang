package main //every go file must have a package name specified

import ( //these are packages imported from the golang standard library
	//formatting strings
	"net/http" //tools for web development

	"github.com/gorilla/mux"
	"github.com/msarah/learngolang/lenslocked.com/views"
)

//Global vars here are temporary
var (
	homeView    *views.View
	contactView *views.View
)

//this will store our template

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil)
	//^^this line will eventually become something like "contactView.render"
	if err != nil {
		panic(err) //panics will clean up later
	}
}

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	//we want to execute the template first before our router

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)

	http.ListenAndServe(":3000", r)
	// starts up a local web server using our new gorilla handler
}
