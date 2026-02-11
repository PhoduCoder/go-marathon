package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	//HandleFunc takes a function with the args as responseWriter and *request
	http.HandleFunc("/cat", c) //Instead of mux.Handle()
	http.HandleFunc("/dog", d)

	http.ListenAndServe(":80", nil) //default serveMux

}

func c(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Cat meows!!")
}

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Dog barks!!")
}
