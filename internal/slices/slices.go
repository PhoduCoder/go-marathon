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

// This is poor since
// you already got the value that you threw away
// and then used the slice lookup again, which was not necessary
func AvoidMax(x []int) int {
	max := x[0]
	for i, _ := range x {
		if x[i] > max {
			max = x[i]
		}
	}
	return max
}

func IdiomaticMax(x []int) int {
	max := x[0]

	for _, v := range x {
		if v > max {
			max = v
		}
	}
	return max
}

func Maximum(x []int) int {
	sort.Ints(x) //Performs an ascending order inplace sort
	return (x[len(x)-1])

}
