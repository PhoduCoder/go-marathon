package main

import "fmt"

func main() {

	//Initialize using a map literal
	m1 := map[int]string{1: "Gaurav", 2: "Saurav", 3: "Raurav"}

	fmt.Println(m1[1])

	//Create map using make
	m := make(map[int]string)
	m[1] = "Shreya"

	fmt.Println(m)

	val, ok := m[1]
	if ok {
		fmt.Println("Value exists")
	} else {
		fmt.Println("Key doesn't exist")
	}
	fmt.Printf("Value is %v ", val)
}
