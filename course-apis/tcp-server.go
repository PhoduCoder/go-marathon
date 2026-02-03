package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

func main() {

	li, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Listener had errors, %v", err)
	}

	defer li.Close()

	var wg sync.WaitGroup
	wg.Add(1)
	go handle(li, &wg)
	wg.Wait()

	fmt.Println("End of program")
}

func handle(li net.Listener, wg *sync.WaitGroup) {
	conn, err := li.Accept()
	if err != nil {
		log.Fatalf("Error when accepting connection %v", err)
	}
	defer wg.Done()
	defer conn.Close()

	io.WriteString(conn, "Hello from TCP server\n")

	//Reading from connection using bufio scanner type
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		fmt.Printf(scanner.Text())
	}
}
