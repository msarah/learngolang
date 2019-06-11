package views

import (
	"html/template"
	"path/filepath"
)

var (
	LayoutDir   string = "views/layouts"
	TemplateExt string = ".gohtml"
)

//NewView returns a View object to render with our specified file
//and any necessary layout files appended.
func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
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

//layoutFiles returns a slice of strings representing
//the layout files used in our application
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	} //panic for now
	return files
}
