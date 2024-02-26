## RW Mutex

The standard library also exposes a [sync.RWMutex](https://pkg.go.dev/sync#RWMutex)

In addition to these methods:

- [Lock()](https://pkg.go.dev/sync#Mutex.Lock)
- [Unlock()](https://pkg.go.dev/sync#Mutex.Unlock)

The `sync.RWMutex` also has these methods:

- [RLock()](https://pkg.go.dev/sync#RWMutex.RLock)
- [RULock()](https://pkg.go.dev/sync#RWMutex.RUnlock)

The `sync.RWMutex` can help with performance if we have a read-intensive
process. Many goroutines can safely read from the map at the same time (multiple
`Rlock()` calls can happen simultaneously). However, only one goroutine can
hold a `Lock()` and all `RLock()`'s will also be excluded.
