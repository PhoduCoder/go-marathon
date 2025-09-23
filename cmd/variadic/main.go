package main

import (
	"fmt"

	"github.com/PhoduCoder/go-marathon/internal/variadic"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	result := variadic.Sum(a...) //Unpacking the slice
	fmt.Println("Sum:", result)
	result2 := variadic.Sum(1, 2, 3, 4, 5) //Directly passing values
	fmt.Println("Sum:", result2)
}
