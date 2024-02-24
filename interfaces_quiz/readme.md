## Interfaces Quiz

Remember, interfaces are collections of method signatures. A type "implements"
an interface if it has all of the methods of the given interface defined on it.

```go
type shape interface {
    area() float64
}
```

If a type in your code implements an `area` method, with the same signature (e.g.
accepts nothing and returns a `float64`), then that object is said to
<em>implement</em> the `shape` interface.

```go
type circle struct {
    radius int
}

// because awe invoke the area() method, we are
// inferring that circle implements shape
func (c *circle) area() float64 {
    return 3.14 * c.radius * c.radius
}
```

This is <em>different from most other languages</em>, where you have to
<em>explicitly</em> assign an interface type to an object like with Java:

```java
class Circle implements Shape
```
