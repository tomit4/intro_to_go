## Slices Review

Slices wrap arrays to give a more general, powerful, and convenient interface to
sequences of data. Except for items with exxplicit dimensions such as
transformation matrices, most array programming in Go is done with slices rather
than simply arrays.

Slices hold references to an underlying array, and if you assign one slice to
another, both refer to the <b>same</b> array. If a function takes a slice
argument, changes it makes to the elements of the slice <em>will be visible to
the caller</em>, analogous to passing a pointer to the underlying array. A Read
function can therefore accept a slice argument rather than a pointer and a
count; the length within the slice sets an upper limit of how much data to read.
Here is the signature of the [Read()](https://pkg.go.dev/os#ReadFile) method of the `File` type in package `OS`:

Referenced from [Effective Go](https://go.dev/doc/effective_go)

```go
func (f *File) Read(buf []byte) (n int, err error)
```
