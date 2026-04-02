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