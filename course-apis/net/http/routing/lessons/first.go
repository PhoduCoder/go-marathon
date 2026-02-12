package main

import (
	"fmt"
	"io"
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
		io.WriteString(w, "My Name is Gaurav")
	})

	http.ListenAndServe(":8080", nil)
}
