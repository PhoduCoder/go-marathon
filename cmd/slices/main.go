package main

import (
	"fmt"

	"github.com/PhoduCoder/go-marathon/internal/slices"
)

func main() {
	num_list := []int{1, 4, 5, 3, 6, 9, 10, 43, 17, 2}
	fmt.Printf("The initial slice of number is %v\n", num_list)

	rev := slices.Reverse(num_list)

	fmt.Printf("The returned reversed slice is %v\n", rev)
}
