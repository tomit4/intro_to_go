## Type Switches

A <em>type switch</em> makes it easy to do several type assertions in a series.

A type switch is similar to a regular switch statement, but the cases specify
<em>types</em> instead of <em>values</em>.

```go
func printNumericValue(num interface{}) {
    switch v := num.(type) {
        case int:
            fmt.Printf("%T\n", v)
        case string:
            fmt.Printf("%T\n", v)
        default:
            fmt.Printf("%T\n", v)
    }
}

func main() {
    printNumericalValue(1)
    // prints "int"

    printNumericalValue("1")
    // prints "string"

    printNumericalValue(struct{}{})
    // prints "struct {}"
}
```

`fmt.Printf("%T\n), v)` prints the <em>type</em> of a variable.
