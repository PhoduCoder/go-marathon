package main

import (
	"log"
	"net/http"

	"github.com/PhoduCoder/go-marathon/internal/echo"
)

func main() {

	log.Fatal(http.ListenAndServe(":8080", echo.Myhandler{}))

}

//go run echo/main.go

//curl localhost:8080 -H "key:value"
