## Slice Comparison in Go
When comparing slices , `==` and `!=` operator are not valid

The comparison of slices can only occur when comparing the entire slice to nil.
So one has two options, 
Either you do compare one by one
OR
Use the reflect package, this provides a DeepCopy() function.

so using `reflect.DeepCopy(slice1, slice2)` can be used to compare