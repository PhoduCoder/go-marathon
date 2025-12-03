package main

import (
	"fmt"
	"sync"
)

var msg string

var wg sync.WaitGroup

func updateMsg(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg = s
}

func main() {
	fmt.Println("Starting main coroutine")
	msg = "Hello World!!"

	wg.Add(3)
	go updateMsg("Hello Grv!", &wg)
	go updateMsg("Hello Gaurav!", &wg)
	go updateMsg("Hello AlienInvader18!", &wg)

	wg.Wait()
	fmt.Println(msg)

	fmt.Println("Ending main coroutine")

}
