package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Starting the main go coroutine")

	var wg sync.WaitGroup

	//MAIN COROUTINE - THINK OF IT AS MANAGER

	wg.Add(3) //Manager thinks i am waiting for three coroutine to finish, puts counter as 3

	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done() //worker says i am doneThe coroutine finishes reducing counter by 1
			fmt.Println("=================Starting Anonymous coroutine=============")
		}()
	}
	//Forking from the main coroutine

	iterations := make([]int, 1)
	for i, _ := range iterations {
		fmt.Println(i)
		fmt.Println("Calling after describing the anonymous coroutine")
	}

	wg.Wait() //Join point - Manager thinks i need to wait until all are done

	fmt.Println("All Coroutines finished")

}
