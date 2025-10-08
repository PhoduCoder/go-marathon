## Split a string 

## Using substring
You can generally use a Split function from the strings package
Note that this function splits on a substring and NOT rune 

```
words := strings.Split(lower_str, " ") //returns a slice of strings
```

Here the separator is " ", blank space

## Split using runes 
Now say you want to split based on your criteria, or a function that you define
that is where you use the FieldsFunc function from the strings package

```
func isSeparator(r rune) bool {
	return r == ',' || r == ';'
}

strings.FieldsFunc(s, isSeparator)

```

This is how the FieldsFunc works. You take a string and pass a function
The function should have the below signature

```
func(r rune) bool
```

The other way to use it is by defining the function inline , in this case, one can use anonymous function 

```
words := strings.FieldsFunc(s, func(r rune) bool {
    return !unicode.IsLetter(r)
})
```


