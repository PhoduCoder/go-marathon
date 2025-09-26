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