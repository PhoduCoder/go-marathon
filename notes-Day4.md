There is no Exception handling in Go, like in Java or python

Instead errors are treated as values, returned by functions and they should be handled

Say if a function returns a division value, 
it will return two values - a) Value obtained after division
b) Error, which will be **nil** in all other cases, except for the case in which the divisor is 0 


Error is actually and interface which can be implemented 
by a) Implementing the error() method 
b) Returning a string from the error() method 

One can declare new error variables using err.New

```
var (
     ErrHourlyRate = errors.New("invalid hourly rate")
     ErrHoursWorked = errors.New("invalid hours worked per week")
)
```

Now these errors ErrHourlyRate and ErrHoursWorked can be passed back 
in the form of errors

When the program executes successfully, i.e. no error,
then we return nil as err. 

This is why it is idiomatic to use 

```
if err != nil :
  fmt.Println("Error is :", err)
```

=======

Panic in Go 




====
Unpack operator in Go

Variadic function and unpack operator 
Variadic functions are used for functions when the number of values of the SAME TYPE 
being passed as ARGUMENT to a function is unknown 

One can use multiple argument types in a function that uses variadic param,
but the variadic param must be the last in the function argument 

It is very similar to *args in Python 

Eg,

```

func sum (i ...int) int {
    total := 0

    for _, num := range(i){
        total += num
    }
    return total
}
```

```
package main
import "fmt"

i := []int{5,10,15,20}
fmt.Println(sum(i...)) #Unpacking the slice of numbers to the variadic function
fmt.Println(sum(5,5,5))
```
