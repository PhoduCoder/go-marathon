package main

import "fmt"

func main() {

	intStream := make(chan int, 4)

	go func() {
		defer close(intStream)
		defer fmt.Println("Producer Done.")
		for i := 0; i < 5; i++ {
			fmt.Println("Sending:", i)
			intStream <- i
			fmt.Println("SENT!!:", i)
		}
	}()

	for v := range intStream {
		fmt.Println("Received", v)
	}
}
