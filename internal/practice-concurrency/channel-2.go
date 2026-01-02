package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Buffered channels output")

	ch := make(chan string, 2)
	ch <- "hello world"
	ch <- "Hello Gaurav"
	close(ch)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println(v)
		}
	}()

	wg.Wait()

	fmt.Println("End of program")

}
