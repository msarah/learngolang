package views

//new package for views

import "html/template"

//make sure to use html/template and not text/template

/*
NewView
creates a new View object
parses all the template files necessary
and returns the new View to us
*/
func NewView(files ...string) *View {
	//Takes one OR MORE files (variadic parameter)

	files = append(files, "views/layouts/footer.gohtml")
	/*
	  Take the files and append existing layout files we want to use
	  This is harcoded now but will change later to choose based on
	  what is in our layout folder
	*/

	//make sure to check for errors
	t, err := template.ParseFiles(files...)
	//unpack slice and parse each value individually (variadic)

	if err != nil {
		panic(err)

	}

	return &View{ //return pointer to View
		Template: t, //our new template with layouts appended
	}
}

type View struct {
	Template *template.Template
	//this simply contains a template
}
