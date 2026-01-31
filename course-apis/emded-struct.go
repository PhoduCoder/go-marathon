package main

import "fmt"

type Person struct {
	fname string
	lname string
	age   int
	//parent_name []string
}

type NRI struct {
	//p Person //1

	Person // This is an example of embedding
	// This causes person attributes to be promoted to NRI's attribute
	// Otherwise had we done `p Person`, we would have to access
	// n.p.fname instead of n.fname
	country_residence string
}

// type National interface {
// 	CurrentAge()
// }

func (p Person) CurrentAge() {
	fmt.Printf("Person %s is of %d age\n", p.fname, p.age)
}

func (n NRI) CurrentAge() {
	// 1
	//fmt.Printf("NRI %s is of %d age\n", n.p.fname, n.p.age)

	fmt.Printf("NRI %s is of %d age\n", n.fname, n.age)
}
func main() {

	p := Person{"Mohan", "Pyare", 25}
	fmt.Println(p)
	p.CurrentAge()

	//g := Person{"Gaurav", "Parashar", 35}
	//n := NRI{g, "USA"}

	n := NRI{Person{"Gaurav", "Parashar", 35}, "USA"}
	fmt.Println(n)
	n.CurrentAge()
	n.Person.CurrentAge()

	//n.p.CurrentAge() // 1 Access the innner type CurrentAge

}
