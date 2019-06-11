package views

<<<<<<< HEAD
import (
	"html/template"
	"net/http"
	"path/filepath"
)

=======
//new package for views

import (
	"html/template"
	"path/filepath"
)

//make sure to use html/template and not text/template

>>>>>>> master
var (
	LayoutDir   string = "views/layouts/"
	TemplateExt string = ".gohtml"
)

<<<<<<< HEAD
//NewView returns a View object to render with our specified file
//and appends any necessary layout files
func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
=======
/*
NewView
creates a new View object
parses all the template files necessary
and returns the new View to us
*/
func NewView(layout string, files ...string) *View {
	//Takes one OR MORE files (variadic parameter)

	files = append(files, layoutFiles()...)
	/*
	  Take the files and append existing layout files we want to use
	  This is harcoded now but will change later to choose based on
	  what is in our layout folder
	*/

	//make sure to check for errors
	t, err := template.ParseFiles(files...)
	//unpack slice and parse each value individually (variadic)
>>>>>>> master
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

type View struct {
	Template *template.Template
	//this simply contains a template
	Layout string
}

<<<<<<< HEAD
//layoutFiles returns a slice of strings representing
//the layout files used in our application
=======
>>>>>>> master
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
<<<<<<< HEAD
	} //panic for now
	return files
}

//Render will render the necessary template based on what
//page is being requested (needs to be an exported function
//because it receives data from handler)
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}
=======
	}
	return files
}
>>>>>>> master
