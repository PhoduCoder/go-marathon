/*
In the coroutines-trick.go, we saw interleaving of prints between coroutines
In Modern Go (1.14+) can preempt more aggressively (even inside long-running functions),
so never assume a goroutine will run to completion without interruption unless you explicitly synchronize it.

In this example, we will use an example to avoid interleaving by using lock
*/

package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	var mu sync.Mutex

	for i, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer mu.Unlock() //Defer follows LIFO (last deferred, first executed)
			//

			mu.Lock()
			fmt.Printf("Executing %d \n", i)
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
}
