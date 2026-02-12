package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Index file")
	})

	http.HandleFunc("/dog/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Dog Index called")
	})

	http.HandleFunc("/me/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "text/html") //Otherwise we see plain-text as content type
		io.WriteString(w, "My Name is Gaurav")
		tpl, err := template.ParseFiles("example.gohtml")
		if err != nil {
			log.Fatalln("error parsing template", err)
		}

		err = tpl.ExecuteTemplate(w, "example.gohtml", "Gaurav")
		if err != nil {
			log.Fatalln("error executing template", err)
		}
	})

	http.ListenAndServe(":8080", nil)
}
