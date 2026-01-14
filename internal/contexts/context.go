package main

import (
	"context"
	"fmt"
)

func main() {
	str := "Gaurav"
	ctx := context.TODO() //Create a context
	// The other way is to use context.Background()
	// Both TODO and background are two ways of starting an EMPTY context
	ctx_new := context.WithValue(ctx, "func_name", "main")
	doSomething(ctx_new, str)
}

func doSomething(ctx context.Context, str string) {

	fmt.Printf("Context passed from %s \n", (ctx.Value("func_name")))
	fmt.Printf(" %s Doing something \n", str)
}

/*
To add a new value to a context, use the context.WithValue function in the context package.
The function accepts three parameters: the parent context.Context, the key, and the value.
The parent context is the context to add the value to while preserving all the other information
about the parent context.
*/
