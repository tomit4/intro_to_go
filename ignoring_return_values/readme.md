## Ignoring Return Values

A function can return a value that the caller doesn't carea bout. We can
explicitly ignore variables using an underscore: '\_' character.

For example:

```go
func getPoint() (x int, y int) int {
    return 3, 4
}

// ignore y value
x, _ := getPoint()
```

Even though `getPoint()` returns two values, we can capture the first one and
ignore the second.

**Why Would You Ignore A Return Value?**

There could be many reasons. For example, maybe a function called `getCircle`
returns the center point and the radius, but you really only need the radius for
your calcuation. In that case you would ignore the center point variable.
