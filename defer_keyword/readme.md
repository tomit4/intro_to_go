## Defer

The `defer` keyword is a farily unique feature of Go. It allows a function to be
executed automatically <em>just before</em> its enclosing function returns.

The deferred call's arguments are evaluted immediately, but the function call is
not executed until the surrounding function returns.

Deferred functions are typically used to close database connections, file
handlers and the like.

For example:

```go
// CopyFile copies a file from srcName to dstName on the local filesystem
func CopyFile(dstName, srcName string) (written int64, err error) {
    // Open the source file
    src, err := os.Open(srcName)
    if err != nil {
        return
    }
    // Close the source file when the Copyfile function returns
    defer src.Close()

    // Create the destination file
    dst, err := os.Create(dstName)
    if err != nil {
        return
    }

    // Close the destination file when the Copyfile function returns
    defer dst.Close()

    return io.Copy(dst, src)
}
```

In the above example, the `src.Close()` function is not called until after the
`CopyFile` function was called but immediately before the `CopyFile` function
returns.
