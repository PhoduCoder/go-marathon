Anonymous functions are very much like normal functions except for two things 

- They don't have a name
- They must be called as soon as they are declared

The second point can actually be avoided by assigning anonymous functions 
to a variable and then calling the variable 

```
func main() {
  f := func() {
    fmt.Println("Executing an anonymous function using a variable")
  }
  fmt.Println("Line after anonymous function declaration")
  f()
}
```

Closures

Anonymous functions have access to variables which are at the level of the function declaration. 
This can be used to implement scenarios like closures in Go

```
func main() {
  i := 0
  incrementor := func() int {
    i +=1
    return i
  }
  fmt.Println(incrementor())  // outputs 1 
  fmt.Println(incrementor())  //outputs 2
  i +=10
  fmt.Println(incrementor()) //outputs 13
}
```

In the above example, incrementor is variable assigned to an anonymous function
This anonymous function accepts no arguments and returns int.
Also note that this anonymous function has access to the variable i, which is NOT passed to the 
anonymous function, but declared at the same level as the anonymous function 

---------
One issue with the above approach is that the main program also have access to variable i 
and can change its value as we saw in line i +=10

Say we only want the value of i to change when we call the anonymous function. 
How do we do that?
We can follow the below steps

------------

1. We can define a function in Go. def abc func() int //abc is a function 
2. That function returns an anonymous function // abc returns an anonymous function, the anonymous function returns an integer
3. This function declares a variable(protected data) and the anonymous function gains access to it since it is defined at the same level as anonymous function definition. 
3. We assign the variable to the function defined in Step 1 // var1 := abc
4. Since the function returns an anonymous function, so essentially the var1 becomes an anonymous function
5. Now everytime we call the variable() like `var1()`, we are calling the anonymous function

A common approach like what is describe above is used to increment , decrement counters etc. 

```
func main() {
  increment := incrementor()
  fmt.Println(increment())
  fmt.Println(increment())
}
func incrementor() func() int {
  i := 0
  return func() int {
    i += 1
    return i
  }
}
```

Same concept as above but passing a struct instead of int 

```
package main

import "fmt"

type Counter struct {
	Value int
}

func main() {
	dec := newDecrementor(10, 2) // start=10, step=2

	fmt.Println(dec()) // {8}
	fmt.Println(dec()) // {6}
	fmt.Println(dec()) // {4}
	fmt.Println(dec()) // {2}
}

// returns a function (closure) that decrements and returns a struct snapshot
func newDecrementor(start, step int) func() Counter {
	c := Counter{Value: start}
	return func() Counter {
		c.Value -= step // mutate captured state
		return c        // return updated struct (by value)
	}
}
```