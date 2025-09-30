package slices

import (
	"sort"
)

func Reverse(x []int) []int {
	rev := make([]int, len(x))

	for i, j := 0, len(x)-1; i < len(x); i, j = i+1, j-1 { //note that you could NOT do i++ or j--
		rev[i] = x[j]
	}
	return rev
}

func Maximum(x []int) int {
	sort.Ints(x) //Performs an ascending order inplace sort
	return (x[len(x)-1])

}
