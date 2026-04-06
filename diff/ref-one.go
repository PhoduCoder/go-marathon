package main

import (
	"fmt"
	"reflect"
)

func main() {

	//i := 42

	v := struct{}{}
	v = {key: 36, value:"million"}

	fmt.Printf("Type of variable is %v\n ", reflect.TypeOf(v))
	fmt.Printf("Value of variable is %v", reflect.ValueOf(v))

}
