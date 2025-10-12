package main

import (
	"fmt"

	"github.com/PhoduCoder/go-marathon/internal/wordcount_new"
)

func main() {

	sent := "this is a black sentence, followed by a white sentence. These lines have no meaning, just to demonstrate the split logic!!"

	wordcount := wordcount_new.Wordcount(sent)

	fmt.Print(wordcount)

}
