package main

import (
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("files/*.gohtml"))
}

type Group struct {
	People []Person
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

	p1 := Person{"Gaurav", "Parashar", 36, "lightblue"}
	p2 := Person{"Gautam", "Prasar", 46, "lightcoral"}
	p3 := Person{"Anil", "Ambani", 63, "lightgreen"}

	group1 := Group{People: []Person{p1, p2, p3}}

	tpl.ExecuteTemplate(w, "sept.gohtml", group1)

}
