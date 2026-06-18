# Go Slices: Length, Capacity, and Underlying Array

## 1. What is a slice?

A slice is **not the actual array**.

A slice is a small structure that points to an underlying array.

Think of a slice as:

```go
pointer to starting element
length
capacity
```

Example:

```go
s := []int{2, 3, 5, 7, 11, 13}
```

Underlying array:

```text
index:  0  1  2  3   4   5
value:  2  3  5  7  11  13
```

At this point:

```text
len = 6
cap = 6
```

---

## 2. Length

Length means:

> How many elements are currently visible in the slice.

Example:

```go
s := []int{2, 3, 5, 7, 11, 13}
s = s[:4]
```

Visible slice:

```text
[2 3 5 7]
```

So:

```text
len = 4
```

---

## 3. Capacity

Capacity means:

> How many elements are available from the slice's current starting point to the end of the underlying array.

This is the most important point.

Capacity is **not always the size of the original array**.

Capacity depends on **where the slice starts**.

---

## 4. Example: slicing from the front

```go
s := []int{2, 3, 5, 7, 11, 13}
```

Underlying array:

```text
index:  0  1  2  3   4   5
value:  2  3  5  7  11  13
        ^
        s starts here
```

So:

```text
len = 6
cap = 6
```

Now:

```go
s = s[:0]
```

Visible slice:

```text
[]
```

But slice still starts at index `0`.

So:

```text
len = 0
cap = 6
```

Now:

```go
s = s[:4]
```

Visible slice:

```text
[2 3 5 7]
```

Still starts at index `0`.

So:

```text
len = 4
cap = 6
```

---

## 5. Example: dropping first two values

```go
s = s[2:]
```

Now the slice starts at index `2`.

```text
index:  0  1  2  3   4   5
value:  2  3  5  7  11  13
              ^
              s starts here
```

Visible slice:

```text
[5 7]
```

Length:

```text
len = 2
```

Capacity is counted from index `2` to the end:

```text
indexes available: 2, 3, 4, 5
```

So:

```text
cap = 4
```

Important:

```text
The underlying array did not change.
Only the slice's starting position changed.
```

---

## 6. `make` with length only

```go
a := make([]int, 5)
```

This creates a slice with 5 visible zero values.

```text
[0 0 0 0 0]
```

So:

```text
len = 5
cap = 5
```

---

## 7. `make` with length and capacity

```go
b := make([]int, 0, 5)
```

This creates an underlying array of capacity 5.

```text
index:  0  1  2  3  4
value:  0  0  0  0  0
```

But the slice has length 0.

So visible slice is:

```text
[]
```

And:

```text
len = 0
cap = 5
```

---

## 8. Why is cap of `c` still 5?

```go
b := make([]int, 0, 5)
c := b[:2]
```

Even though `b` has length `0`, it has capacity `5`.

So Go allows slicing up to capacity.

`c` starts at index `0`.

```text
index:  0  1  2  3  4
value:  0  0  0  0  0
        ^
        c starts here
```

Visible slice:

```text
[0 0]
```

So:

```text
len = 2
cap = 5
```

Because capacity is counted from index `0` to the end.

---

## 9. What if we do `c := b[1:2]`?

```go
b := make([]int, 0, 5)
c := b[1:2]
```

Now `c` starts at index `1`.

```text
index:  0  1  2  3  4
value:  0  0  0  0  0
           ^
           c starts here
```

Visible slice:

```text
[0]
```

Length:

```text
len = 2 - 1 = 1
```

Capacity is counted from index `1` to the end:

```text
indexes available: 1, 2, 3, 4
```

So:

```text
cap = 4
```

---

## 10. Main formula

For:

```go
newSlice := oldSlice[low:high]
```

Length is:

```text
high - low
```

Capacity is:

```text
old capacity - low
```

Example:

```go
b := make([]int, 0, 5)
c := b[1:2]
```

So:

```text
len(c) = 2 - 1 = 1
cap(c) = 5 - 1 = 4
```

---

## 11. Reslicing rule

You can reslice up to capacity.

Example:

```go
b := make([]int, 0, 5)
c := b[:2]
```

This is allowed because:

```text
2 <= cap(b)
```

But this is not allowed:

```go
c := b[:6]
```

Because:

```text
6 > cap(b)
```

This will panic.

---

## 12. Why does this panic?

```go
b := make([]int, 0, 5)
c := b[1:2]
d := c[2:5]
```

First:

```go
c := b[1:2]
```

So:

```text
len(c) = 1
cap(c) = 4
```

Now:

```go
d := c[2:5]
```

This tries to slice up to `5`.

But `c` only has capacity `4`.

So this panics:

```text
panic: runtime error: slice bounds out of range [:5] with capacity 4
```

Valid version:

```go
d := c[2:4]
```

---

## 13. Mental model

Always draw the underlying array.

Then ask:

```text
Where does this slice start?
How many elements are visible?       => length
How many elements remain till end?   => capacity
```

Example:

```text
array: [0 0 0 0 0]
index:  0 1 2 3 4
```

For:

```go
c := b[1:2]
```

Slice starts at index `1`.

Visible indexes:

```text
1
```

So:

```text
len = 1
```

Remaining indexes from start:

```text
1, 2, 3, 4
```

So:

```text
cap = 4
```

---

## 14. One-line summary

A slice's capacity is not the total size of the original array.

It is:

> The number of elements available from the slice's current starting position to the end of the underlying array.
