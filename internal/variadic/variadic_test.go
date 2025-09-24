package variadic

import "testing"

func TestSum(t *testing.T) {
	num := []int{1, 2, 3, 4, 5}
	got := Sum(num...)
	want := 15

	if got != want {
		t.Fatalf("Sum() = %d; want %d", got, want)
	}
}
