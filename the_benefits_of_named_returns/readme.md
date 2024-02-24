## The Benefits Of Named Returns

### Good For Documentation (Understanding)

Named return parameters are great for documenting a function. We knnow what the
function is returning directly from its signature, no need for a comment.

Named return parameters are particularly important in longer functions with many
return values.

```go
func calculator(a, b int) (mul, div int, err error) {
    if b == 0 {
        return 0, 0, errors.New("Can't divid by zero")
    }
    mul = a * b
    div = a / b
    return mul, div, nil
}
```
