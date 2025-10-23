Interface vs generics

Say if you wanted to implement this 
```
func PrintInts(nums []int) {
    for _, n := range nums {
        fmt.Println(n)
    }
}

func PrintStrings(words []string) {
    for _, w := range words {
        fmt.Println(w)
    }
}

```

Regular interfaces work when different types implement the same method signature (same input and return types).

If your functions need to accept different input types or return different types, a single interface can’t cover them.

That’s where generics (parametric polymorphism) shine: they let you write one function or type that works for multiple input/output types safely.

So:

Interface → behavioral polymorphism, same method signature.

Generics → parametric polymorphism, flexible types.