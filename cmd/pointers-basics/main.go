package main

//Let us first understand why do we need pointers

import (
	"fmt"
)

func main() {

	var a int = 3
	var b int = 4

	fmt.Println("Initial values of a and b")
	fmt.Println(a, b)

	fmt.Println("Calling increment function")
	incrementInt(a, b)
	fmt.Println("Calling a and b in main after increment function")
	fmt.Println(a, b)

	fmt.Println("=========Pointers=============")

	var p int = 9
	var q int = 10

	fmt.Printf("Initial values of p and q are %d & %d\n", p, q)

	incrementviaPointer(&p, &q) //Now passing the pointer to p and q instead of copy

	fmt.Printf("Values of p and q from main after the increment via pointer function are %d and %d \n", p, q)

}

// increment via pointer function that takes a int pointer
func incrementviaPointer(m, n *int) {

	fmt.Printf("Inside the increment via pointer function values of m and n are %d and %d\n", *m, *n)
	//To dereference the pointer
	// i.e. to get the value stored by the variable whose pointer is represented
	// use *m
	*m = *m + 1
	*n = *n + 1
	fmt.Printf("After increment logic in pointer m and n are %d and %d\n", *m, *n)

}

func incrementInt(x, y int) {
	fmt.Println("Printing passed values")
	fmt.Println(x, y)
	x += 1
	y += 1
	fmt.Println("Inside function after logic printing")
	fmt.Println(x, y)
}

//As you notice in this above example, when calling the increment function
// we passed a copy of values of a and b to increment function.

//When we called the a and b value from the main function,
//after the increment, the value in main remained same
//since we actually passed a copy of values
