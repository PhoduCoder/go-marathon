package main

import (
	"log"
	"net/http"

	"github.com/PhoduCoder/go-marathon/internal/rolldice"
)

func main() {
	http.HandleFunc("/rolldice", rolldice.Rolldice)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
