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

}
