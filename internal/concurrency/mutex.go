package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.RWMutex

var counter int = 10

func increment(counter *int, wg *sync.WaitGroup) {
	mu.Lock()
	*counter++
	fmt.Println("Running increment counter")
	fmt.Println("New value of counter is", *counter)
	defer mu.Unlock()
	defer wg.Done()
}

func decrement(counter *int, wg *sync.WaitGroup) {
	mu.Lock()
	*counter--
	fmt.Println("Running decrement counter")
	fmt.Println("New value of counter is", *counter)
	defer mu.Unlock()
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
