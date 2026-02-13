package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

type index int

// type dog struct{}
type me struct{}

var c index

// var d dog
var m me

func (c index) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Index file")
}

// Using HandlerFunc
// So we didn't define dog struct and a variable, we can simply substitute any function with that signature (w http.ResponseWriter, req *http.Request)
func dogwrite(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Dog Index called")
}

func (m me) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html") //Otherwise we see plain-text as content type

	//w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, "My Name is Shreya")
	tpl, err := template.ParseFiles("example.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}

	err = tpl.ExecuteTemplate(w, "example.gohtml", "Shreya")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func main() {

	http.Handle("/", c)

	http.Handle("/dog/", http.HandlerFunc(dogwrite))

	http.Handle("/me/", m)

	http.ListenAndServe(":8080", nil)
}
