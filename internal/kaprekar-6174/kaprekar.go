package kaprekar6174

import "fmt"
import "sort"

func Kaprekar6174(n int) (int, error) {

	// take a 4 digit number as input 
	// rearrange the digits to form the largest and smallest numbers possible
	// subtract the largest from smallest 
	// Get the resultant and continue the process until we reach 6174
	// count the number of iterations to reach 6174
	// finally publish the number and the iteration count 

	for i:=1; i <=10; i++ {
		largest := largestNumber(n)
		smallest := smallestNumber(n)
		result := largest - smallest
		if result == 6174 {
			return i, nil
		} else {
			kaprekar6174(result)
		}
	}

func largestNumber(n int) int {
	digits := make([]int, 4)
	for i := 0; i < 4; i++ {
		digits[i] = n % 10
		n /= 10
	}
	fmt.Println(digits)

	fmt.Println(sort.Ints(digits))

}


