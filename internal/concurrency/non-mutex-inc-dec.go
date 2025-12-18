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

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		if (i % 3) == 0 {
			go decrement(&counter, &wg)
		} else {
			go increment(&counter, &wg)
		}
	}
	// wg.Add(3)
	// go decrement(&counter, &wg)
	// go increment(&counter, &wg)
	// go increment(&counter, &wg)

	wg.Wait()

	fmt.Println("Final value of counter is", counter)
	fmt.Println("===========**********============")

}

//Important Note
// While running the code locally may feel like
// there is no data race, this can go wrong at scale

//Also it can lead to inconsistent results
// since coroutines run non-deterministic fashion

//counter = 10
// G1 reads 10 → increments to 11
// G2 reads 10 → decrements to 9
// Final value = 9 or 11 (never 10)

// counter++ is three operations:

// Load from memory

// Add 1

// Store back to memory

// A mutex makes those 3 steps atomic as a unit.
