package main

import (
	"fmt"

	"github.com/PhoduCoder/go-marathon/internal/controlflow"
)

func main() {
	test := controlflow.Fizzbuzz(43)
	fmt.Printf("Test is %v", test)
}
