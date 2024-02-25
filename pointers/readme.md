## Pointers

Pointers hold the memory address of a value.

The `*` sytnax defines a pointer.

```go
var p *int
```

A pointer's zero value is `nil`

The `&` operator generates a pointer to its operand.

```go
myString := "hello"
myStringPtr = &myString
```

The `*` dereferences a pointer to gain access to the value

```go
fmt.Println(*myStringPtr) // read myString through the pointer
*myStringPtr = "world" // set myString through the pointer
```

Unlike C, Go has no [pointer arithmetic](https://en.wikipedia.org/wiki/Pointer_arithmetic)

### Just Because You Can Doesn't Mean You Should

We're doing this exercise to understand that pointers <b>can</b> be used in this
way. That said, pointers can be <em>very</em> dangerous. It's generally a better
idea to have your functions accept non-pointers and return new values rather
than mutating pointer inputs.
