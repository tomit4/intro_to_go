## Buffered Channels

Channels can <em>optionally</em> be buffered.

### Creating A Channel With A Buffer

You can provide a buffer length as the second argument to `make()` to create a
buffered channel:

```go
ch := make(chan int, 100)
```

Sending on a buffered channel only blocks when the buffer is <em>full</em>.

Receiving blocks only when the buffer is <em>empty</em>.
