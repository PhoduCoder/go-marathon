package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Started the main coroutine")

	var wg sync.WaitGroup
	//var stream chan string //declare a string channel

	stream := make(chan string)
	wg.Add(1)

	go func() {
		defer wg.Done()
		close(stream)
		//No point in writing to a closed channel
		//stream <- "Hello World" //Writing to the channel

		fmt.Println("From the coroutine")
	}()

	msg, ok := <-stream

	fmt.Println("Values are ", msg, ok) // Value changes to false. since channel is closed

	//fmt.Println(<-stream)
	wg.Wait()
	fmt.Println("Exiting the coroutine")
}
