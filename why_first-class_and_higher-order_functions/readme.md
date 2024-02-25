## Why First-Class And Higher-Order Functions?

At first, it may seem like dynamically creating functions and passing them
around as variables adds unnecessary complexity. Most of the time you would be
right. There are cases however when functions as values makea lot of sense. Some
of these include:

- [HTTP API](https://pkg.go.dev/net/http) handlers
- [Pub/Sub](https://pkg.go.dev/cloud.google.com/go/pubsub) handlers
- Onclick callbacks

Any time you need to run custom code at <em>a time in the future</em>, functions
as values might make sense.

### Definition: First-Class Functions

A first-class function is a function that can be treated like any other value.
Go supports firts-class functions. A function's type is dependent on the types
of its parameters and return values. For example, these are differnt function
types:

```go
func() int
```

```go
func(string) int
```

### Definition: Higher-Order Functions

A higher-order function is a function that takes a function as an argument or
returns a function as a return value. Go supports higher-order functions. For
example, this function takes a function as an argument:

```go
func aggregate(a, b, c int, arithmetic func(int, int) int) int {}
```
