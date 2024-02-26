## Closing Channels In Go

Channels can be explicitly closed by a <em>sender</em>:

```go
ch := make(chan int)

// do some stuff with the Channels

close(ch)
```

### Checking If A Channel Is Closed

Similar to the `ok` value when accessing data in a map, receivers can check if
the `ok` value when receiving from a channel to test if a channel was closed.

```go
v, ok := <-ch
```

ok is `false` if the channel is empty and closed.

### Don't Send On A Closed Channel

Sending on a closed channel will cause a panic. A panic on the main goroutine
will cause the entire program to crash, and a panic in any other goroutine will
cause <em>that goroutine</em> to crash.

Closing isn't necessary. There's nothing wrong with leaving channels open,
they'll still be garbage collected if they're unused. You should close channels
to indicate explicitly to a receiver that nothing else is going to come across.
