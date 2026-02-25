package main

import (
	"net/http"
)

func main() {

	//http.Handle("/", http.FileServer(http.Dir("assets")))
	//http.HandleFunc("/img", profile)
	http.ListenAndServe(":80", http.FileServer(http.Dir("assets")))
}

// func profile(w http.ResponseWriter, req *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	io.WriteString(w, `<img src="adventure.jpg">`)
// }

//Note that if you have a file named index.html then
// that file will be served
// and you won't see the rest of directory
