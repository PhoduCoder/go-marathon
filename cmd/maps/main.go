package main

import (
	"fmt"

	"github.com/PhoduCoder/go-marathon/internal/maps"
)

func main() {
	users := maps.GetUsers()
	fmt.Printf("The users are %v map\n", users)
	fmt.Printf("The first user is %s\n", users[1])
}
