package main

import "fmt"

func main() {

	type User struct {
		name    string
		age     int
		subject string
	}

	//Creating maps of different types
	//remember values are of any kind

	m1 := make(map[int]User)

	m1[1] = User{"Gaurav", 36, "Physics"}
	m1[2] = User{"Shreya", 33, "Geography"}
	m1[3] = User{"Goli", 24, "Biology"}
	m1[4] = User{"Ramba-HO", 12, "History"}

	for ind, val := range m1 {
		fmt.Printf("Entry %d => name is %v is %d years old and favorite subject is %v \n", ind, val.name, val.age, val.subject)
	}
}
