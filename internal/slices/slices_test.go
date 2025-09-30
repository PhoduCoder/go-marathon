package slices

import (
	"reflect"
	"testing"
)

func TestSlices(t *testing.T) {
	got := Reverse([]int{1, 2, 3, 4, 5})

	want := []int{5, 4, 3, 2, 1}

	// if got != want {
	// 	t.Fatalf("Reverse() = %q; want %q", got, want)
	// }

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Reverse() = %q; want := %q", got, want)
	}

}
