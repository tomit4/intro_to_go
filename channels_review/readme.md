## Channels Review

Here are a few extra things you should understand about channels from [Dave
Cheney's awesome article](https://dave.cheney.net/2014/03/19/channel-axioms).

### A Send To A Nil Channel Blocks Forever

```go
var c chan string // c is Nil
c <- "let's get started" // blocks
```

### A Receive From A Nil Channel Blocks Forever

```go
var c chan string // c is nil
fmt.Println(<-c) // blocks
```

### A Send To A Closed Channel Panics

```go
var c = make(chan int, 100)
close(c)
c <- 1 // panic: send on a close channel
```

### A Receive From A Closed Channel Returns The Zero Value Immediately

```go
var c = make(chan int, 100)
close(c)
fmt.Println(<-c) // 0
```
