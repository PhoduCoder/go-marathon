/*
Send 5 integers from main to a goroutine and print them.
Use a buffered channel (size 3). Observe when sends block.
Close a channel and use range to read until done.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)

	go func() {
		time.Sleep(5 * time.Second)
		for v := range ch {
			fmt.Printf("Receiving %d element value\n", v)
			fmt.Println(v)
		}
	}()

	for i := 0; i < 6; i++ {
		fmt.Printf("Sending value %d to channel\n", i)
		ch <- i
		fmt.Printf("Sent value %d to channel\n", i)

	}
	close(ch)

}
