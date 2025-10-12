package wordcount_new

import (
	"fmt"
	"strings"
)

func Wordcount(sent string) map[string]int {

	//Separate sentence by spaces and punctuation marks
	words := strings.FieldsFunc(sent, separator) //returns []string

	fmt.Println(words)

	fmt.Println(len(words))

	wordcount := make(map[string]int)

	for _, value := range words {
		wordcount[value]++
	}

	return wordcount

}

func separator(r rune) bool {

	return r == ' ' || r == ',' || r == '.' || r == '!'

}
