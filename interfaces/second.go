package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct{}
type Cat struct{}

func (d Dog) Speak() {
	fmt.Println("woof")
}

func (c Cat) Speak() {
	fmt.Println("meow")
}

func main() {
	var s Speaker

	s = Dog{}
	s.Speak()

	s = Cat{}
	s.Speak()
}

// JUST FOR UNDERSTANDING
// REAL WORLD YOUU WONT DO INTERFACE VARIABLE EQUAL TO STRUCT
