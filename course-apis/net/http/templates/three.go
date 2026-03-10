package main

import (
	"net/http"
	"os"
	"text/template"
)

func main() {

	tpl := template.Must(template.ParseGlob("files/*.gohtml"))

	//names := []string{"Gaurav", "shreya"}

	colors := map[string]string{"Gaurav": "lightblue", "Shreya": "lightred"}

	tpl.ExecuteTemplate(os.Stdout, "five.html", colors)

	http.ListenAndServe(":80", http.FileServer(http.Dir("files")))

}
