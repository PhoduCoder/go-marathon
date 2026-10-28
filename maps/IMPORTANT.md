Type category	                                        Behavior with var
Value types (int, bool, struct, array)	                Fully usable immediately
Reference types (map, slice, pointer, channel)	        Nil until initialized, make helps it to initialize

So make is very helpful since it initializes the container not necessaarily the content 

```
m := make(map[string]int)
```
This makes map usaable buut it can't create entries with zero vaalues 
since we dont know what keys would be

With slices it is a little different 

```
s := make([]int, 3) //Initialize a slice with length and capacity as 3 
fmt.Println(s) // [0 0 0]
```

make(map) → “create empty dictionary” 
make([]T, n) → “create array of size n filled with zero values” 


A Go map has two states worth distinguishing:

```
var cache1 map[string]string // nil map
cache = make(map[string]string) // initialized, empty map
```

Both have length zero, and both are fine to read from:
But only the initialized map can be written to. Assigning into a nil map panics:

A nil map has not been initialized. You can safely read from it; a missing key returns the value type’s zero value. You can also check whether a key exists with the two-value lookup. But you cannot add entries to a nil map.

A map created with make is initialized and ready for entries. Both maps can be empty, but only the initialized map can be written to.


```
package main

import "fmt"

func main() {

	var cache map[string]string

	cache1 := make(map[string]string)

	fmt.Println("From make one")
	fmt.Println(len(cache1))
	fmt.Println(cache1["app"] == "")
	fmt.Println(cache1 == nil)
	fmt.Println("This declares a map and initialized with zero values, so in case of string it is \"\" ")

	fmt.Println("===============")
	fmt.Println("From var initialized map")
	fmt.Println(len(cache))
	fmt.Println(cache["app"] == "")
	fmt.Println(cache == nil)
	fmt.Println("This declares a nil map")
}

```


## Reference Like type 
A Go map is a reference-like type.
When you pass a map to a function, Go copies the map value, but that copy still refers to the same underlying map data.

Copy contains a reference to the same underlying map data.

```
func change(m map[string]string) {
    m["app"] = "Grafana"
}

func main() {
    cache := make(map[string]string)
    cache["app"] = "ArgoCD"

    change(cache)

    fmt.Println(cache["app"])
}
```
