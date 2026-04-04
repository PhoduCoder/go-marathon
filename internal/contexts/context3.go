package main

import (
	"context"
	"fmt"
)

// meaningful data
type RequestMeta struct {
	UserID  string
	TraceID string
}

// unexported key type
type contextKey struct{}

var metaKey = contextKey{}

func main() {
	ctx := context.Background()

	meta := &RequestMeta{
		UserID:  "user-123",
		TraceID: "trace-xyz",
	}

	str := "Gaurav"

	// put into context
	ctx = context.WithValue(ctx, metaKey, meta)

	doSomething(ctx, str)
}

func doSomething(ctx context.Context, str string) {
	// retrieve
	meta, ok := ctx.Value(metaKey).(*RequestMeta)
	if !ok {
		fmt.Println("no metadata found")
		return
	}

	fmt.Printf("UserID: %s\n", meta.UserID)
	fmt.Printf("TraceID: %s\n", meta.TraceID)

	fmt.Printf("Passed argument is %s \n", str)
}
