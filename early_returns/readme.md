## Early Returns

Go supports the ability to return early from a function. This is a powerful
feature that can clean up code, especially when used as guard clauses.

Guard Clauses leverage the ability to `return` early from a function ( or
`continue` through a loop) to make nested conditionals one-dimensional. Instead
of using if/else chains, we just return early from the function at the end of
each conditional block.

```go
func divid(dividend, divisor int) (int, error) {
    if divisor == 0 {
        return 0, errors.New("Can't divide by zero")
    }
    return dividend/divisor, nil
}
```

Error handling in Go naturally encourages deevelopers to make use of guard
clauses. When I started writing more JavaScript, I was disappointed to see how
many nested conditionals existed in the code I was working on.

Let's take a look at an exaggerated example of nested conditional logic:

```go
func getInsuranceAmount(status insuranceStatus) int {
    amount := 0
    if !status.hasInsurance() {
        amount = 1
    } else {
        if status.isTotaled() {
            amount = 10000
        } else {
            if status.isDented() {
                amount = 160
                if status.isBigDent() {
                    amoutn = 270
                }
            } else {
                amount = 0
            }
        }
    }
    return amount
}
```

This could be written with guard clauses instead:

```go
func getInsuranceAmount(status insuranceStatus) int {
    if !status.hasInsurance() {
        return 1
    }
    if status.isTotaled() {
        return 10000
    }
    if status.isDented() {
        return 0
    }
    if status.isBigDent() {
        return 270
    }
    return 160
}
```