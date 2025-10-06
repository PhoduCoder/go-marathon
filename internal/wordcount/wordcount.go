package wordcount

import (
	"fmt"
	"strings"
)

func WordCount(str string) map[string]int {
	//Function takes a sentence
	//returns the count of each word in the sentence
	//map[string]int

	//str := "Ram was the greatest god, Hanuman was his best friend who fought by his side"

	lower_str := strings.ToLower(str)

	fmt.Println("Lowered case sentence is %v", lower_str)

	words := strings.Split(lower_str, " ") //returns a slice of strings

	//fmt.Println(words)

	//return words
	m := make(map[string]int)

	for _, word := range words {
		//one can replace the entire below logic
		//from 29-34 in one line `m[word]++`
		val, ok := m[word]
		if !ok {
			m[word] = 1 //When key doesn't exist in map, then assign a value of 1
		} else {
			m[word] = val + 1 //When key exist then increment that value
		}

	}
	return m

}
