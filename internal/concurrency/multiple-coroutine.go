//This is a Go coroutine with multiple coroutines
// Objective is to show that how multiple coroutine takes
// way lesser time to complete
//Compared to single which takes 5 seconds, this completes in 500 ms
//So approximately 10 times faster

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)         //This is like incrementing the counter
		go heavyTask(&wg) // Starts a coroutine - Pay attention to go keyword
		//wg is passed to make sure that the coroutine can call wg.Done()

	}
	wg.Wait() // This is the main coroutine waiting for all coroutines to finish
	fmt.Println("Total time: ", time.Since(start))
}

func heavyTask(wg *sync.WaitGroup) {
	defer wg.Done() //This is decrementing the counter in the coroutine, notice the defer
	time.Sleep(500 * time.Millisecond)
	fmt.Println("heavy task completed")
}
