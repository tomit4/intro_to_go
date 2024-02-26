## Channels

Channels are a typed, thread-safe queue. Channels allow different goroutines to
communicate with each other.

### Creat A Channel

Like maps and slices, channels must be created <em>before</em> use. They also
use the same `make` keyword:

```go
ch := make(chan int)
```

### Send Data To A Channel

```go
ch <- 69
```

The `<-` operator is called the <em>channel operator</em>. Data flows in the
direction of the arrow. This operation will <em>block</em> until another
goroutine is ready to receive the value.

### Receive Data From A Channel

```go
v := <-ch
```

This reads and removes a value from the channel and saves it into the variable
`v`. This operation will <em>block</em> until there is a value in the channel to
be read.

### Blocking And Deadlocks

A [deadlock](https://programming.guide/go/detect-deadlock.html) is when a group
of goroutines are all blocking so none of them can continue. This is a common
bug that you need to watch out for in concurrent programming.
