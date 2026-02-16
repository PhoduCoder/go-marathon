package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/img", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html")

	io.WriteString(w, `<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">`)
}

//Loading from an external URL
