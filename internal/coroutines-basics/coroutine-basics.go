package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Starting the main go coroutine")

	var wg sync.WaitGroup

	wg.Add(1) //Adding a counter with the number of coroutines

	go func() {
		defer wg.Done()
		fmt.Println("=================Starting Anonymous coroutine=============")
	}() //Forking from the main coroutine

	wg.Wait() //Join point

	iterations := make([]int, 5)
	for i, _ := range iterations {
		fmt.Println(i)
		fmt.Println("Calling after describing the anonymous coroutine")
	}

}
