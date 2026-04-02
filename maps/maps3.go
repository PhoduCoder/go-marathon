package main

import "fmt"

func main() {

	age := map[string]int{"Gaurav": 36, "Shreya": 33, "Goli": 24, "Rumba-HO": 12}

	//Note that the order of iteration is NOT guaranteed
	for name, val := range age {
		fmt.Printf("%v is %d years old\n", name, val)
	}
}
