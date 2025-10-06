package main

import (
	"fmt"

	"github.com/PhoduCoder/go-marathon/internal/wordcount"
)

func main() {

	sentence := "Gaurav is a good boy and an active learner, gaurav is gritty and hard working, he is an example for generations to come"

	words := wordcount.WordCount(sentence)

	fmt.Println(words)
}
