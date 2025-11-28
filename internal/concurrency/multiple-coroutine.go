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
		wg.Add(1)
		go heavyTask(&wg)

	}
	wg.Wait()
	fmt.Println("Total time: ", time.Since(start))
}

func heavyTask(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(500 * time.Millisecond)
	fmt.Println("heavy task completed")
}
