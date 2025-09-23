package main

import (
	"fmt"

	"github.com/PhoduCoder/go-marathon/internal/divide"
)

func main() {
	result, err := divide.Divide(10, 20)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
