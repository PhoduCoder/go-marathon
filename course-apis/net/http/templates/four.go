package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("files/*.gohtml"))
}

func main() {

	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	colors := map[string]string{
		"Gaurav": "lightblue",
		"Shreya": "lightcoral",
	}

	tpl.ExecuteTemplate(w, "five.gohtml", colors)
}
