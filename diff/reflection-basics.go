// Example program to show difference between
// Type and Kind and to demonstrate use of
// Methods provided by Go reflect Package
package main

import (
	"fmt"
	"reflect"
)

type address struct {
	houseNo int
	street  string
}

type details struct {
	address
	fname   string
	lname   string
	age     int
	balance float64
}

type myType string

func showDetails(i, j interface{}) {
	t1 := reflect.TypeOf(i) // This gives main.details
	k1 := t1.Kind()         //This will give struct
	t2 := reflect.TypeOf(j) //This gives main.myType
	k2 := t2.Kind()         // This gives string
	fmt.Println("Type of first interface:", t1)
	fmt.Println("Kind of first interface:", k1)
	fmt.Println("Type of second interface:", t2)
	fmt.Println("Kind of second interface:", k2)

	fmt.Println("The values in the first argument are :")
	if reflect.ValueOf(i).Kind() == reflect.Struct {
		value := reflect.ValueOf(i)
		fmt.Printf("%T type and value %v\n", value, value) // gives this reflect.Value type and {Go Geek 32 33000.203}

		numberOfFields := value.NumField() // returns the number of fields present in a struct, panics if type is not struct

		// NumField(): This method returns the number of fields present in a struct. If the passed argument is not of the Kind reflect.Struct then it panics.
		// Field(): This method allows us to access each field in the struct using an Indexing variable.

		for i := 0; i < numberOfFields; i++ {
			fmt.Printf("%d.Type:%T || Value:%#v\n",
				(i + 1), value.Field(i), value.Field(i))

			fmt.Println("Kind is ", value.Field(i).Kind())
		}
	}
	value := reflect.ValueOf(j)
	fmt.Printf("The Value passed in "+
		"second parameter is %#v\n", value)
}

func main() {
	iD := myType("12345678")
	person := details{
		address: address{310, "Tenth street"},
		fname:   "Go",
		lname:   "Geek",
		age:     32,
		balance: 33000.203,
	}
	showDetails(person, iD)
}

// Either all fields named
// Or all positional
// Never mix

//full named

// person := details{
// 	address: address{
// 		houseNo: 310,
// 		street:  "Tenth street",
// 	},
// 	fname:   "Go",
// 	lname:   "Geek",
// 	age:     32,
// 	balance: 33000.203,
// }

//Full positional
// person := details{
// 	address{310, "Tenth street"},
// 	"Go",
// 	"Geek",
// 	32,
// 	33000.203,
// }
