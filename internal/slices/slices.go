package slices

func Reverse(x []int) []int {
	rev := make([]int, len(x))

	for i, j := 0, len(x)-1; i < len(x); i, j = i+1, j-1 {
		rev[i] = x[j]
	}
	return rev
}
