package main

import (
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var testTemplate *template.Template

type User struct {
	Name          string
	Uname         string //has to be capitalized in order to be exported
	Age           int
	FavouriteFood map[string]string
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	data := User{
		"Sarah Murphy",
		"msarah",
		22,
		map[string]string{
			// "Breakfast": "Oatmeal",
			"Lunch":  "Sandwich",
			"Dinner": "Pasta",
		},
	}

	err := testTemplate.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {

	var err error
	testTemplate, err = template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	err = testTemplate.Execute(os.Stdout, User{
		"Sarah Murphy",
		"msarah",
		22,
		map[string]string{
			"Breakfast": "Oatmeal",
			"Lunch":     "Sandwich",
			"Dinner":    "Pasta",
		},
	})
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	http.ListenAndServe(":8080", r)
}

// data.Name = "Sarah"
//
// err = t.Execute(os.Stdout, data)
// //arg1 where you want it to spit the data out
// //arg2 data to be transmitted
// if err != nil {
// 	panic(err)
// }
