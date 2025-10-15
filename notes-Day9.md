## Pointers in Go

When you call a function in Go, depending on the variables types
passed to the function, you would mostly pass a copy of values

So if you don't return from the called function and reassign the values
you won't notice data changing

Look at examples under cmd/pointers-basics


However the case of slices is a little different.
Slices itself is a struct that points to an array.

A slice in Go is already a lightweight descriptor (it contains a pointer to an underlying array, plus length and capacity):

So when you pass a slice to a function, Go copies this small header, but both headers point to the same underlying array.

```
func modify(s []int) {
    s[0] = 99 // changes underlying array
}

func main() {
    nums := []int{1, 2, 3}
    modify(nums)
    fmt.Println(nums) // [99 2 3]
}
```

But if you reassign the slice, then the caller won't see the change
this is because it now becomes a new slice 

```
func appendValue(s []int) {
    s = append(s, 4)
}

func main() {
    nums := []int{1, 2, 3}
    appendValue(nums)
    fmt.Println(nums) // [1 2 3]
}
```

So to now make the above work, you use a pointer 

```
func appendValue(s *[]int) {
    *s = append(*s, 4)
}

func main() {
    nums := []int{1, 2, 3}
    appendValue(&nums)
    fmt.Println(nums) // [1 2 3 4]
}

```

It is idiomatic to use a pointer receiver for a method that modifies a slice.

```
type Ints []int

func (s *Ints) Add(x int) {
    *s = append(*s, x)
}

nums := Ints{1, 2, 3}
nums.Add(4)
```

Instead of 

```
func Add(s *[]int, x int) {
    *s = append(*s, x)
}

nums := []int{1, 2, 3}
Add(&nums, 4)
```

also use the pointer receiver when using structs 

```
type Person struct {
    Name string
    Age  int
}

func (p *Person) Birthday() {
    p.Age++ // modifies the original
}

person := Person{"Alice", 30}
person.Birthday()
fmt.Println(person.Age) // 31
```

If the method doesn't modify the struct, then one can use the value receiver instead
```
func (p Person) Greet() {
    fmt.Println("Hello,", p.Name)
}
```

Also keep in mind that when defining the pointer receiver over a struct 
calling changes 

```
type Person struct {
    Name string
    Age  int
}

func (p *Person) Birthday() {
    p.Age++
}

func main() {
    person := Person{"Alice", 30}

    person.Birthday()  // ✅ works, even though receiver is *Person
    (&person).Birthday() // also works explicitly

    fmt.Println(person.Age) // 31
}
```

## Calling the function when using slices and pointer receiver

```

type Ints []int

func (s *Ints) Add(x int) {
    *s = append(*s, x)  // modify the original slice
}


func main() {
    nums := Ints{1, 2, 3}

    nums.Add(4)      // ✅ works, Go automatically takes the address
    (&nums).Add(5)   // also works explicitly

    fmt.Println(nums) // [1 2 3 4 5]
}

```

Contrasting this with a function not using pointer receiver 

```
func AddElement(s *[]int, x int) {
    *s = append(*s, x)
}

func main() {
    nums := []int{1, 2, 3}
    AddElement(&nums, 4)  // must pass &nums
    fmt.Println(nums)      // [1 2 3 4]
}
```