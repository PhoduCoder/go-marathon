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
	//mux := http.NewServeMux() //Returns a pointer to mux
	//WE DON"T CREATE THE MUX, RATHER USE THE DEFAULTMUX
	//THIS IS DONE BY PASSING NIL IN LISTENANDSERVE

	http.Handle("/cat", c) //Instead of mux.Handle()
	http.Handle("/dog", d)

	http.ListenAndServe(":80", nil)

}

func (c cat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Cat meows!!")
}

func (d dog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Dog barks!!")
}
