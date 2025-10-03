package main

import (
	"fmt"

	"github.com/PhoduCoder/go-marathon/internal/maps"
)

func main() {
	users := maps.GetUsers()
	fmt.Printf("The users are %v map\n", users)
	fmt.Printf("The first user is %s\n", users[1])
	users[5] = "Swapnil"
	fmt.Println("After modification users map", users)

	user7, err := (users[7])
	fmt.Printf("Does User7 exist ,decision is %v, it was supposed to be %s\n", err, user7)

	users[7] = "Sonal"
	user7, err = (users[7])
	fmt.Printf("Does User7 exist ,decision is %v, it was supposed to be %s\n", err, user7)
	fmt.Println("After modification users map", users)

}
