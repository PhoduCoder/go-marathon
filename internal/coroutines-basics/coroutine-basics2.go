package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Starting the main go coroutine")

	var wg sync.WaitGroup

	//MAIN COROUTINE - THINK OF IT AS MANAGER

	anon := func() {
		defer wg.Done() //worker says i am doneThe coroutine finishes reducing counter by 1
		fmt.Println("=================Starting Anonymous coroutine=============")
	} //Forking from the main coroutine

	wg.Add(1) // Manager thinks i am waiting for one coroutine to finish, puts counter as 1

	go anon() //Calling the anonymous function, but appending 'go' makes it coroutine

	iterations := make([]int, 1)
	for i, _ := range iterations {
		fmt.Println(i)
		fmt.Println("Calling after describing the anonymous coroutine")
	}

	wg.Wait() //Join point - Manager thinks i need to wait until all are done

	fmt.Println("All Coroutines finished")

}
