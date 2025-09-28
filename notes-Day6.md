Function Types

Functions can be assigned to variables
Functions can be returned inside a function
Functions can also be passed as an argument to a function 

We can also define function types in Go 

Function type signature = Types of input params + Return values 

----
```
package main

import "fmt"

type calc func(int, int) string // Defines a new type called calc

//Any function that has the same signature, i.e. takes two integer inputs
// and returns a string is of the type calc

//But you might wonder why do something like this

func add(i, j int) string {
	return fmt.Sprintf("Added %d + %d = %d", i, j, i+j)
}

func mult(a, b int) string {
	return fmt.Sprintf("Multiplying two numbers %d and %d = %d", a, b, (a * b))
}

//Now one can define a function that takes calc type

func Calculator(f calc, a, b int) string {
	return f(a, b)
}

func main() {
	sum := Calculator(add, 5, 10)
	mult := Calculator(mult, 5, 10)

	fmt.Printf("The sum and product are %s and %s\n", sum, mult)
}
```