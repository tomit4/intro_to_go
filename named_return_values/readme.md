## Named Return Values

Return values may be given names, and if they are, then theey are treated the
same as if they were new variables defined at the top of the function.

Named return values are best thought of as a way to document the purpose of the
returned values.

According to the [tour of go](https://go.dev/tour/welcome/1):

> A return statement without arguments returns the named return values. This is
> known as a "naked" return. Naked return statements should be used only in
> short functions. They can harm readability in longer functions.

```go
func getCoords() (x, y int) {
    // x and y are initialized with zero values

    return // automatically returns x and y
}
```

Is the same as:

```go
func getCoords() (int, int) {
    var x int
    var y int
    return x, y
}
```

In the first example, `x` and `y` are the return values. At the end of the
fucnction, we could simply write `return` to return the values of those two
variables, rather than writing `return x,y`.
