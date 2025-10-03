package wordcount

import (
	"fmt"
	"strings"
)

func WordCount(str string) map[string]int {
	//Function takes a sentence
	//returns the count of each word in the sentence

	str := "Ram was the greatest god, Hanuman was his best friend who fought by his side"

	lower_str := strings.ToLower(str)

	fmt.Println("Lowered case sentence is %v", lower_str)

}
