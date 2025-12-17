//Both these programs don't use mutex and still
// try to update a shared resource
// This is a classic case of data race

// One can check it by running `go run -race non-mutex-inc-dec.go`

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

var counter int = 10

func increment(counter *int, wg *sync.WaitGroup) {
	*counter++
	fmt.Println("Running increment counter")
	fmt.Println("New value of counter is", *counter)
	defer wg.Done()
}

func decrement(counter *int, wg *sync.WaitGroup) {
	*counter--
	fmt.Println("Running decrement counter")
	fmt.Println("New value of counter is", *counter)
	defer wg.Done()
}

func main() {

	fmt.Println("Initial Value of counter is", counter)

	wg.Add(3)
	go decrement(&counter, &wg)
	go increment(&counter, &wg)
	go increment(&counter, &wg)

	wg.Wait()

	fmt.Println("Final value of counter is", counter)
	fmt.Println("===========**********============")

}
