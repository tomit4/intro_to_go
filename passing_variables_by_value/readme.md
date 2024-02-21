## Passing Variables By Value

Variabes in Go are passed by value (expcept for a few data types we haven't
covered yet). "Pass by value" means that when a variable is passed into
a function, that function receives a <em>copy</em> of the variable. The function
is unable to muatate the caller's data.

```go
func main() {
    x := 5
    increment(x)

    fmt.Println(x)
    // still prints 5,
    // because the increment function
    // received a copy of x and NEVER RETURNED IT
}

func increment(x int) {
    x++
}
```

A better way to implement this basic functionality would be to return the x
value from the `increment()` function, <em>and</em> re-assign x to that return
value like so:

```go
func main() {
    x := 5
    x = increment(x)

    fmt.Println(x)
}

func increment(x int) {
    x++
    return x
}
```
