## Anonymous Structs In Go

An anonymous struct is just like a normal struct, but it is defined without a
name and therefore cannot be refrerenced elsewhere in the code.

To create an anonymous struct, just instantiate the instance immediately using a
second pair of brackets after declaring the type:

```go
myCar := struct {
    Make string
    Model string
} {
    Make: "tesla",
    Model: "model 3"
}
```

You can even nest anonymous structs as fields within other structs:

```go
type car struct {
    Make string
    Model string
    Height int
    Width int
    // Wheel is a field containing an anonymous struct
    Wheel struct {
        Radius int
        Material string
    }
}
```

## When Should You Use An Anonymous Struct?

In general, <em>prefer named structs</em>. Named structs make it easier to read
and understand your code, and they have the nice side-effect of being reusable.
I sometimes use anonymous structs when I <em>know</em> I won't ever need to use
a struct again. For example, sometimes I'll use one to create the shape of some
JSON data in HTTP handlers.

If a struct is only meant to be used once, then it makes sense to declare it in
such a way that developers down teh road won't be tempted to accidentally use it
again.
