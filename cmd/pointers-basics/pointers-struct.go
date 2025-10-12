package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Salary float64
}

func main() {
	fmt.Println("Pointers with Strings, slices and structs")

	p1 := Person{"Gaurav", 35, 10000000.00}

	fmt.Printf("Initial Salary is %f\n", p1.Salary)

	increment(p1)
	fmt.Printf("Salary after increment without pointers is %f\n", p1.Salary)
	//Here we pass a pointer
	incrementSalary(&p1)
	fmt.Printf("Salary after increment is %f\n", p1.Salary)
}

func incrementSalary(p *Person) {
	p.Salary = 2000000.50 //Note that in structs we don't deference like in int
	//so we don't do *p.Salary, rather that is understood for structs

	//p.Salary is syntactic sugar for (*p).Salary

}

func increment(p Person) {
	p.Salary = 3000000.00
}
