In Go, else must appear on the same line as the closing brace } of the preceding if block.

so the below is invalid

```
if x%3 == 0 {
    fmt.Println("Fizz")
}
else if x%5 == 0 {  // <- error
    fmt.Println("Buzz")
}
```

while the following is valid

```
if x%3 == 0 {
    fmt.Println("Fizz")
} else if x%5 == 0 {
    fmt.Println("Buzz")
}
```