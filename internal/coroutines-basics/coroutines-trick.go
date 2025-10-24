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
			fmt.Printf("Executing %d \n", i)
			defer wg.Done()
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
