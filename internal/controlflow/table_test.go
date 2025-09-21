package controlflow

import "testing"

func TestFizzbuzz_Table(t *testing.T) {
	//Defining a struct of test cases
	cases := []struct {
		name string
		n    int
		want []string
	}{
		{"n=0", 1, []string{"0", "1"}},
		{"n=5", 5, []string{"0", "1", "2", "Fizz", "4", "Buzz"}},
		{"n=15", 15, []string{"0", "1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
	}
	for _, tc := range cases {
		//Using t.Run to run subtests
		//The first argument is the name of the subtest
		//The second argument is a function that takes a *testing.T parameter
		//This function contains the actual test logic
		//This allows us to run each test case independently and get separate results for each
		t.Run(tc.name, func(t *testing.T) {
			got := Fizzbuzz(tc.n)
			if len(got) != len(tc.want) {
				t.Fatalf("len(got)=%d len(want)=%d", len(got), len(tc.want))
			}
			for i := range tc.want {
				if got[i] != tc.want[i] {
					t.Fatalf("i=%d got=%q want=%q", i, got[i], tc.want[i])
				}
			}
		})
	}
}
