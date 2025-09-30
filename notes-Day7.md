## Slice Comparison in Go
When comparing slices , `==` and `!=` operator are not valid

The comparison of slices can only occur when comparing the entire slice to nil.
So one has two options, 
Either you do compare one by one
OR
Use the reflect package, this provides a DeepCopy() function.

so using `reflect.DeepCopy(slice1, slice2)` can be used to compare

##Range function for Slice

```
z := []int{100,200,300,500}
for index, value := range(z){
    fmt.Println(index,value) //will get 0,100 1,200 ...
}
```

```
	word := "GaurAv"
	for index, value := range word {
		fmt.Println(index, value)
	}
```
Outputs 
0 71 [Unicode code point for 'G']
1 97 [Unicode for 'a']
2 117['u']
3 114 ['r']
4 65 [Unicode for 'A']
5 118 ['v']

###DataTypes for Range function
Basically range function works for 
slices [Gives index and value]
Strings [Gives index and rune]
Maps [Gives key and value]
Channels [Gives value(one at a time until channel is closed)]