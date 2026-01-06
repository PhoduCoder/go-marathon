package main

import "fmt"

func main() {

	//This creates an anoymous function that returns a reader only channel
	// Inside this anonymous function is where
	// we define the channel
	// we also define the coroutine that writes to the coroutine
	channel_var := func() <-chan int {
		ch := make(chan int, 5) //Initialize the channel

		//Coroutine that writes to the channel
		go func() {
			defer close(ch)
			for i := 10; i < 12; i++ {
				//Writing to the channel
				ch <- i
			}
		}()
		return ch // we return a channel that is populated async by a goroutine
	}

	// all of initialize the channel, write to channel & close the channel
	// is exposed via a reader channel, it is an unidirectional channel
	channel_reader := channel_var()

	for j := range channel_reader {
		fmt.Println(j)
	}
}

/*
Keep in mind that this is done so that
we don't run into panic issues when dealing
with writing and reading of channel
*/
