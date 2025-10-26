package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	for i, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("Executing %d \n", i)
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
}

/*
If you want all goroutines to see the correct value for their iteration, you must pass it as an argument:

for _, s := range []string{"hello","greetings","good day"} {
    go func(x string) {
        fmt.Println(x)
    }(s) // Passing it as an argument
}
*/

/*
OUTPUT 

======
xecuting 2 
good day
Executing 1 
Executing 0 
greetings
hello
(base) MacBookPro:go-marathon alieninvader$ go run internal/coroutines-basics/coroutines-trick.go 
Executing 2 
good day
Executing 1 
Executing 0 
hello
greetings
==========
*/

/*
If you notice the output, you will realise that go coroutines 1 and 0 were 
preempted by the Go runtime and hence the output appeared as 
Executing 1, Executing 0, hello and greetings instead of 
Executing 1, greetings, Executing 0 and  hello

This is because Each goroutine does two separate prints (Printf then Println).
The goroutine can be preempted between those calls, so other goroutines run and print in between — causing the mixed output order you observed.