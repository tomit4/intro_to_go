## Anonymous Functions

Anonymous functions are true to form in that they have <em>no name</em>. We've
been using them throughout this chapter, but we haven't really talked about them
yet.

Anonymous functions are useful when defining a function that will only be used
once or to create a quick [closure](https://golangdocs.com/closures-in-golang).

```go
// doMath accepts a functioin that converts one int into another
// and slice of ints. It returns a clice of ints that have been
// converted by the passed in function.
func doMath(f func(int) int, values []int) []int {
    var result []int
    for _, n := range nums {
        result = append(result, f(n))
    }
    return result
}

func main() {
    nums := []int{1, 2, 3, 4, 5}

    // Here we define an anonymous function that doubles an int
    // and pass it to doMath
    allNumsDoubled := doMath(func(n int) int {
        return n * 2
    }, nums)

    fmt.Println(allNumsDoubled)
    // prints:
    // [2 4 6 8 10]
}
```
