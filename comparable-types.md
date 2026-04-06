Can Go compare two values of this type using built-in == without ambiguity?

If yes, it is comparable.

If no, it is not.


# Comparable
bool
numbers
string
pointers
channels
arrays of comparable elements
structs with all comparable fields
interfaces whose dynamic values are comparable
custom types based on comparable underlying types

# Not comparable
slices
maps
functions
structs containing any non-comparable field
arrays containing non-comparable elements