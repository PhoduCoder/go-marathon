package main

import (
	"fmt"
	"log"
	"net/http"
)

type Person int // Need to define this globally
// so that both functions have access to this

func main() {

	var p Person

	http.ListenAndServe(":8080", p)

}

func (p Person) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalf("Fatal error occurred when parsing form")
	}
	host := req.Host
	remote := req.RemoteAddr
	fmt.Println("Calling host is ", host)
	fmt.Println("Calling remote is ", remote)

	url := req.Form

	for key, val := range url {
		fmt.Println("Form is", key, val)
	}

	//Getting form value for a field say `name`
	name := req.FormValue("name")
	fmt.Printf("Name passed is", name)

	fmt.Fprintf(resp, "Writing person age here")

}

//CURL REQUEST
//curl http://localhost:8080?name=shreya -H "mykey:myvalue" -d "name":"gaurav" -d "key":"value"
