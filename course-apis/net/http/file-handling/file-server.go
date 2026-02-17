package main

import (
	"io"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("assets")))
	http.HandleFunc("/img", profile)
	http.ListenAndServe(":80", nil)
}

func profile(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, `<img src="adventure.jpg">`)
}
