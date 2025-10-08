package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Hello, 世界")

	//Note how multi line strings is declared in Go
	my_sentence := `Ram is a good boy 世界, but he is bound to be good 世界, as he is an incarnation of Lord Vishnu! However Ram lived as a simple human bound by humanly limitation
	and his life shows 世界 how should one live`

	list := strings.FieldsFunc(my_sentence, func(r rune) bool {
		//The below removes the 世界 since 世界 is considered a letter but in the below you are actually only taking english alphabet characters
		//return r < 'A' || (r > 'Z' && r < 'a') || r > 'z' // [Ram is a good boy but he is bound to be good as he is an incarnation of Lord Vishnu However Ram lived as a simple human bound by humanly limitation and his life shows how should one live]
		return !unicode.IsLetter(r) // [Ram is a good boy 世界 but he is bound to be good 世界 as he is an incarnation of Lord Vishnu However Ram lived as a simple human bound by humanly limitation and his life shows 世界 how should one live]
	})

	fmt.Println(list)
}
