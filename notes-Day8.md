## Split a string 

## Using substring
You can generally use a Split function from the strings package
Note that this function splits on a substring and NOT rune - equivalent of char in GO 

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


rune literal   → ' '   // single quotes → one character
string literal → " "   // double quotes → string type

Always use single quotes ' ' when comparing runes.


So FieldsFunc is taking a function here which actually returns true
for when fields are ! and ,. Note the '' and not "", since later would be a string in GO

```
fields := strings.FieldsFunc(s, func(r rune) bool {
	return r == '!' || r == ','
})
```

-----
FieldsFunc internally loops through the string.
Each rune is passed to your function f.
If your function returns true, Go treats that rune as a split point (separator).
So the logic of returning true is with you. 

So the function itself doesn’t split anything — it just tells Go where to split.

You can think of it like this:

“Whenever f(r) returns true, start a new word after this character.”
-----

This is powerful for single character rules 

----

say we have a string with many punctuations, however we want to ignore all punctuations 
and only get the words

```
s := "Hello, world! Go: is great."
	
	words := strings.FieldsFunc(s, func(r rune) bool {
		// Split on anything that is NOT a letter, see the exclamation before the unicode function
		return !unicode.IsLetter(r)
	})

	fmt.Println(words) // returns this [Hello world Go is great]
```

