package main

import (
	"fmt"
	"sync"
)

func main() {
	var rw sync.RWMutex
	var wg sync.WaitGroup
	var data int

	for i := 0; i < 6; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			rw.RLock()
			fmt.Printf("Coroutine %d running reading data value %d\n", id, data)
			fmt.Printf("Interleaving test\n")
			rw.RUnlock()
		}(i)
	}

	go func() {
		wg.Add(1)
		defer wg.Done()
		rw.Lock()
		data++
		fmt.Printf("Write coroutine running with data %d \n", data)
		rw.Unlock()
	}()

	wg.Wait()
	fmt.Println("Main function running")
}
