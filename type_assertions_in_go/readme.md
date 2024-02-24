## Type Assertions In Go

When working with interfaces in Go, every once-in-awhile you'll need access to
the underlying type of an interface value. You can cast an interface to its
underlyinig type using a <em>type assertion</em>.

```go
type shape interface {
    area() float64
}

type circle struct {
    radius float64
}

// "c" is a new circle cast frojm "s"
// which is an instance of a shape.
// "ok" is a bool that is true if s was a circle
// or false if s isn't a circle
c, ok := s.(circle)
```
