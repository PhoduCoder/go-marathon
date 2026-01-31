//YOU CAN DEFINE A METHOD THAT TAKES AN INTERFACE

//BUT YOU CAN'T DEFINE A METHOD FOR AN INTERFACE
// like func (h Human)speakAge(){}

// SO basically a function that can take many different types
// and can change it behavior based on what it takes

// LIKE makeSound function
// Can take different animal types and make different noises

package main

import "fmt"

type Person struct {
	fname string
	lname string
	age   int
}

type NRI struct {
	Person
	country_residence string
}

// You defined an interface
// which also has method definition
// You can't define the method on the interface
// func (h Human)speakAge(){} WONT work, since you are defining a method on interface
type Human interface {
	CurrentAge()
}

// Defining a method that takes an interface
// YOU CANNOT DEFINE A NEW METHOD ON INTERFACE
// func (h Human)speakAge(){} WONT work, since you are defining a method on interface

func speakAge(h Human) {
	h.CurrentAge()
}

// Person implements Human Interface
func (p Person) CurrentAge() {
	fmt.Printf("Person %s is of %d age\n", p.fname, p.age)
}

// NRI implements human interface
func (n NRI) CurrentAge() {
	fmt.Printf("NRI %s is of %d age\n", n.fname, n.age)
}
func main() {

	p := Person{"Mohan", "Pyare", 25}
	//fmt.Println(p)
	speakAge(p)

	n := NRI{Person{"Gaurav", "Parashar", 35}, "USA"}
	//.Println(n)
	speakAge(n)

}
