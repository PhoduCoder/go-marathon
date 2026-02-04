package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {

	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error occured when listening, %v \n", err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalf("Error occurred when accepting connection %v \n", err)
		}
		go handle(conn)

	}

}

func handle(conn net.Conn) {

	err := conn.SetDeadline(time.Now().Add(20 * time.Second))
	if err != nil {
		log.Fatalf("Connection timeout, possibly deadline expired %s", err)
	}

	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		//NOTE that fmt.Printf(conn, "your string") WONT work, since fmt.Printf takes a string and not io.Reader()
		fmt.Fprintf(conn, "Reply from the server, we got your request and are working on it\n")
	}

	//Note that this line never executes since there is a for loop scanner.Scan()
	// This preludes the fact that we comment line 39-41
	//and that scanner.Scan() doesn't end until the client closes the connection
	// or an error happened
	//One can try with telnet and once we exit out of telnet client
	// One would see the below line getting executed
	fmt.Printf("End of coroutine accepting connection\n")

}
