package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/img", me)
	http.HandleFunc("/me.jpg", myphoto)
	http.ListenAndServe(":80", nil)
}

func myphoto(w http.ResponseWriter, req *http.Request) {
	file, err := os.Open("me.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		log.Fatalf("file not found!!!!")
	}
	defer file.Close()

	//Copy the content of file to writer
	// Since file implements Reader Interface
	io.Copy(w, file)
}

func me(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html")
	io.WriteString(w, `<img src="/me.jpg">`)
}
