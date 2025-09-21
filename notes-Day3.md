Notes about testing

Three RULES 
[ unit tests file naming convention]
[ Unit test function name has to start with Test capitalized T]
[Unit test function signature ]

For whitebox unit tests, the following rules hold
a) The file must be named origfile_test.go. As an example
say your logic had control.go,  so your test file name becomes 
control_test.go

b) The function must start with Testwhatever
So func signature must be like 
```
func Testcontrol(t *testing.T) {}
```

c) this function doesn't return anything, only contains t *testing.T as 
the argument 
======

Notes about declaring list of strings 

a) var a []string
  a = {"grv", "test", "go"}

b) var list []string 
list = append(list, "apple")
list = append(list, "banana")

c) list := []string{"apple","banana"}

d) list := make([]string, 3)
list[0] = "apple"
list[1] = "banana"
list[2] = "cherry"

Always remember that list of strings will have double quotes,
you can't have single quotes

Single quotes represent rune literals
which are an similar to characters 

====

String formatting 

%q is for strings or slices of strings
%d is for int values 
%v is the general one, should work for most types 
%c is for runes 
