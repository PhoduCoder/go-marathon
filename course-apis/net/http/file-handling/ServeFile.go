package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/me", profile)
	http.HandleFunc("/push.jpg", mypic)
	http.ListenAndServe(":80", nil)
}

func mypic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "push.jpg")

}

func profile(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-type", "text/html")
	io.WriteString(w, `<img src="/push.jpg">`)
}
