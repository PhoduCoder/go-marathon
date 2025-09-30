package main

import (
	"fmt"

	"github.com/PhoduCoder/go-marathon/internal/slices"
)

func main() {
	num_list := []int{1, 4, 5, 3, 6, 9, 110, 43, 17, 2}
	fmt.Printf("The initial slice of number is %v\n", num_list)

	new_list := []int{1, 2, 3, 4, -1, -4, 7, 10, 8, 9}

	rev := slices.Reverse(num_list)

	fmt.Printf("The returned reversed slice is %v\n", rev)

	fmt.Printf("The maximum number in original slice is %d\n", slices.AvoidMax(num_list))

	fmt.Printf("The maximum number in original slice via idiomatic Go is %d\n", slices.IdiomaticMax(new_list))

	fmt.Printf("The maximum number in the slice is %d\n", slices.Maximum(new_list))

	fmt.Printf("Printing the slice %v & reverse %v again\n", num_list, rev) //Note that the slice is sorted now because of inplace sorting
}
