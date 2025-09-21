package controlflow

import "testing"

func TestFizzbuzz(t *testing.T) {
	got := Fizzbuzz(15)
	want := []string{"0", "1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}

	// if len(got) != len(want) {
	// 	t.Fatalf("Fizzbuzz() = %q; want %q", got, want)
	// }
	for i := range got {
		if got[i] != want[i] {
			t.Fatalf("Fizzbuzz() = %q; want %q", got, want)
		}
	}
}

func TestLenFizzBuzz(t *testing.T) {
	got := Fizzbuzz(15)
	want := []string{"0", "1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}

	if len(got) != len(want) {
		t.Fatalf("FizzBuzz() = %d; want %d", len(got), len(want))
	}
}
