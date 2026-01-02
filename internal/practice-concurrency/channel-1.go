package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("This is the main coroutine")

	//declaring a buffered channel
	ch := make(chan int)

	go func() {
		start := time.Now()
		fmt.Println("Running from a coroutine")
		//Created a coroutine and is now blocking
		receive := <-ch
		fmt.Printf("Received value from channel is", receive, "\n")
		end := time.Since(start)
		fmt.Println("Time taken was ", end)
	}()

	//Sending a number to channel and will block until some one is receiving on the channel
	ch <- 3
	fmt.Println("Sleeping for 5 seconds")
	time.Sleep(5 * time.Second)
	defer close(ch)
}
