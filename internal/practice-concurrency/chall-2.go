package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	done := make(chan struct{})

	go func() {
		for v := range ch {
			fmt.Println("Receiving", v)
		}
		done <- struct{}{}
	}()

	for i := 0; i < 6; i++ {
		fmt.Println("Sending", i)
		ch <- i
	}

	close(ch)
	<-done // wait for receiver to finish
}
