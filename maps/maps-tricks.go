package main

import "fmt"

func main() {

	var m map[string]int //This declares m, but does NOT allocate memory for the map.
	//So internally m = nil, there is no underlying hash table

	m["x"] = 1 // HENCE panic

	fmt.Println(m["x"]) // This works and returns 0 since GO allows reading from a nil map

}
