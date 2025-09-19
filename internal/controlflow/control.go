package controlflow

import (
	"strconv"
)

func Fizzbuzz(n int) []string {

	list := make([]string, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			list[i] = "0"
		} else if i%3 == 0 && i%5 == 0 {
			list[i] = "FizzBuzz"
		} else if i%3 == 0 {
			list[i] = "Fizz"
		} else if i%5 == 0 {
			list[i] = "Buzz"
		} else {
			list[i] = strconv.Itoa(i)
		}
	}
	return list
}

// func main() {
// 	test := Fizzbuzz(34)
// 	fmt.Printf("Test is %v", test)
// }
