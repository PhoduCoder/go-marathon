package main

import (
	"context"
	"fmt"
)

type RequestMeta struct {
	UserID  string
	TraceID string
}

func main() {

	str := "Gaurav"

	meta := RequestMeta{
		UserID:  "123",
		TraceID: "abc-xyz",
	}

	ctx_new := context.WithValue(context.Background(), "meta", meta)
	doSomething(ctx_new, str)
}

func doSomething(ctx context.Context, str string) {

	// Retrieve
	fmt.Printf("Context passed from %+v \n", (ctx.Value("meta").(RequestMeta)))
	fmt.Printf(" %s Doing something \n", str)
}
