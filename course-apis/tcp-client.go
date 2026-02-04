package main

import (
	"fmt"
	"io"
	"net"
)

func main() {

	//Dialing as a client to a tcp server
	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		fmt.Println("Error dialing connection")
	}
	defer conn.Close()

	//since ioutil.ReadAll accepts an io reader
	bs, err := io.ReadAll(conn)
	if err != nil {
		fmt.Printf("Error when reading from the dialed connection %v ", err)
	}
	fmt.Println(string(bs))
}
