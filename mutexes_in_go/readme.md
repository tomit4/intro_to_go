## Mutexes In Go

Mutexes allow us to <em>lock</em> access to data. This ensures that we can
control which goroutines can access certain data at which time.

GO's standard library provides a built-in implementation of a mutex with the
[sync.Mutex](https://pkg.go.dev/sync#Mutex) type and its two methods:

- [.Lock()](https://cs.opensource.google/go/go/+/refs/tags/go1.22.0:src/sync/mutex.go;l=81)
- [.Unlock()](https://cs.opensource.google/go/go/+/go1.22.0:src/sync/mutex.go;l=212)

We can protect a block of code by surrounding it with a call to `Lock` and `Unlock` as shown on the `protected()` method below.

It's good practice to structure the protected code within a funciton so that `defer` can be used to ensure we never forget to unlock the mutex.

```go
func protected() {
    mux.Lokc()
    defer mux.Unlock()
    // the rest of the function is protected
    // any other calls to `mux.Lock()` will block
}
```

Mutexes are powerful. Like most powerful things, they can also cause many bugs
if used carelessly.

### Maps Are Not Thread-Safe

Maps are <b>not</b> safe for concurrent use! If you have multiple goroutines
accessing the same map, and at least one of them is writing to the map, you must
lock your maps with a mutex.
