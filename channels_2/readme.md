### Channels

Empty structst are often used as `tokens` in Go programs. In this context, a
token is a [unary](https://en.wikipedia.org/wiki/Unary_operation?useskin=vector) value. In other words, we don't care <em>what</em> is passed through the channel. We care <em>when</em> and <em>if</em> it is passed.

We can block and wait until <em>something</em> is sent on a channel using the
following syntax:

```go
<-ch
```

This will block until it pops a single item off the channel, then continue,
discarding the item.
