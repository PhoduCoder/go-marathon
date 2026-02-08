package main

import (
	"fmt"
	"io"
	"net/http"
)

type cat string

type dog int

func main() {

	var c cat
	var d dog

	//Create a new mux
	mux := http.NewServeMux() //Returns a pointer to mux
	mux.Handle("/cat", c)
	mux.Handle("/dog", d)

	http.ListenAndServe(":80", mux)

}

func (c cat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Cat meows!!")
}

func (d dog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Dog barks!!")
}
