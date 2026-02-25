package echo

import (
	"fmt"
	"net/http"
)

//First define a type that can implement the handler interface(serveHTTP method)
// of the http package

type Myhandler struct{}

func (h Myhandler) HandleFunc("/chapter1", func(w http.ResponseWriter, r *http.Request) {
    msg := "Chapter 1"
    w.Write([]byte(msg))
})

func (h Myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := "<h1>Hello Gaurav</h1>"
	fmt.Println(r.Header, r.Body)

	w.Write([]byte(msg))

}
