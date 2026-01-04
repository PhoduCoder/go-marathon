package main

import "fmt"

func main() {

	ch := make(chan int)
	ch1 := make(chan int, 3)

	ch1 <- 3
	ch1 <- 6
	ch1 <- 9

	close(ch)
	close(ch1)

	buff_count := 0
	unbuff_count := 0

	for i := 0; i < 500; i++ {

		select {
		case <-ch:
			//fmt.Println("Unbuffered selected")
			unbuff_count++

		case <-ch1:
			//fmt.Println("Buffered channel")
			buff_count++
		}
	}

	fmt.Printf("Unbuffered and buffered are invoked %d and %d times respectively", unbuff_count, buff_count)

}
