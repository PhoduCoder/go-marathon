package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok %s\n", time.Now().Format(time.RFC3339))
	})

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       60 * time.Second, // server kills idle keep-alive conns after 60s
	}

	log.Println("server listening on http://127.0.0.1:8080 (IdleTimeout=60s)")
	log.Fatal(srv.ListenAndServe())
}
