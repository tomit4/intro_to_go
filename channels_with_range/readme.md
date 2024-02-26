## Range

Similar to slices and maps, channels can be ranged over.

```go
for item := range ch {
    // item is the next value received from the channel
}
```

This example will receive values over the channel (blockinb at each iteration if
nothing new is there) and will exit only when the channel is closed.
