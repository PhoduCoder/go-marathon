package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct {
	name string
}

func (d Dog) Speak() {
	fmt.Println(d.name, "says woof")
}

func main() {
	d := Dog{name: "rocky"}

	var s Speaker
	s = d
	s.Speak()
}

// JUST FOR UNDERSTANDING
// REAL WORLD YOUU WONT DO INTERFACE VARIABLE EQUAL TO STRUCT
