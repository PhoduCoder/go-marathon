package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct {
	name string
}

type Cat struct {
	name string
}

func (c Cat) Speak() {
	fmt.Println(c.name, "says meow")
}

func (d Dog) Speak() {
	fmt.Println(d.name, "says woof")
}

func makeItSpeak(s Speaker) {
	s.Speak()
}

func main() {
	d := Dog{name: "rocky"}
	makeItSpeak(d)

	c := Cat{name: "Milo"}
	makeItSpeak(c)
}
