package main

import (
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("files/*.gohtml"))
}

type Person struct {
	//Fields have to be capitalized??
	Fname string
	Lname string
	Age   int
	Color string
}

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, req *http.Request) {

	p1 := Person{"Gaurav", "Parashar", 36, "lightcoral"}

	tpl.ExecuteTemplate(w, "six.gohtml", p1)

}
