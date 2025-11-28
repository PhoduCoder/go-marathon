package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	for i := 0; i < 10; i++ {
		heavyTask()

	}

	fmt.Println("Total time: ", time.Since(start))
}

func heavyTask() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("heavy task completed")
}
