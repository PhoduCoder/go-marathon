package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start main")
	doWork()
	fmt.Println("end main")
}

func doWork() {
	fmt.Println("start doWork")

	for i := 0; i < 3; i++ {
		fmt.Println("step", i)
		time.Sleep(50 * time.Millisecond)
	}

	fmt.Println("end doWork")
}
