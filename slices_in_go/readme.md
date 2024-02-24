## Slices In Go

<em>99 times out of 100</em> you will use a slice instead of an array when
working with ordered lists.

Arrays are fixed in size. Onceyou make an array like `[10]int` you can't add an
11th element.

A slice is a <em>dynamically-szized, flexible</em> view of the elements of an
array.

Slices <b>always</b> have an underlying array, though it isn't always specified
explicitly. To explicitly create a slice on top of an array we can do:

```go
primes := [6]int{2, 3, 5, 7, 11, 13}
mySlices := primes[1:4]
// mySliceds = {3, 5, 7}
```

The syntax is:

```
arrayname[lowIndex:highIndex]
arrayname[lowIndex:]
arrayname[:highIndex]
arrayname[:]
```

Where `lowIndex` is inclusive and `highIndex` is exclusive

Either `lowIndex` or `highIndex` or both can be omitted to use the entire array
on that side.
