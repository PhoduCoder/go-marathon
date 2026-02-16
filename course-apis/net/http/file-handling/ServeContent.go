package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/me", profile)
	http.HandleFunc("/push.jpg", mypic)
	http.ListenAndServe(":80", nil)
}

func mypic(w http.ResponseWriter, req *http.Request) {
	file, err := os.Open("push.jpg")
	if err != nil {
		http.Error(w, "File doesnot exist", 404)
		log.Fatalf("This file is not FOUND!!")
	}

	defer file.Close()

	fstat, err := file.Stat()
	if err != nil {
		http.Error(w, "File doesnot exist", 404)
		log.Fatalf("This file is not FOUND!!")
	}

	http.ServeContent(w, req, file.Name(), fstat.ModTime(), file)

}

func profile(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-type", "text/html")
	io.WriteString(w, `<img src="/push.jpg">`)
}
