Pointer Receivers

A receiver type on a method can be a pointer.

Methods with pointer receivers can modify the value to which the receiver
points. Since methods often neeed to modify their receiver, pointer receivers
are <em>more common</em> than value receivers.

### Pointer Receiver

```go
type car struct {
    color string
}

func (c * car) setColor(color string) {
    c.color = color
}

func main() {
    c := car{
        color: "white",
    }
    c.setColor("blue")
    fmt.Println(c.color)
    // prints "blue"
}
```

### Non-Pointer Receiver

```go
type car struct {
    color string
}

func (c car) setColor(color string) {
    c.color = color
}

func main() {
    c := car{
        color: "white",
    }
    c.setColor("blue")
    fmt.Println(c.color)
    // prints "white"
}
```
