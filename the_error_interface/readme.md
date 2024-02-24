## The Error Instance

Go programs express errors with `error` values. An Error is any type that
implements the simple built-in [error interface](https://go.dev/blog/error-handling-and-go):

```go
type error interface {
    Error() string
}
```

When something can go wrong in a function, that function should return an `error` as its last return value. Any code that calls a function that can return an `error` should handle errors by testing whether the error is `nil`.

```go
// Atoi converts a stringified number to any integer
i, err := strconv.Atoi("42b")
if err != nil {
    fmt.Println("couldn't convert:", err)
        // becuase "42b" isn't a valid integer, we print:
        // couldn't convert: strconv.Atoi: parsing "42b": invalid syntax
        // Note:
        // 'parsing "42b": invalid syntax' is returned by the .Error() method
        return
}
// if we get here, then
// i was converted successfully
```

A `nil` error denotes success; a non-nil error denotes failure.

## The Error Interface

Because errors are just interfaces, you can build your own custom types that
implement the `errror` interface. Here's an example of a `userError` struct that
inplements the `error` interface:

```go
type userError struct {
    name string
}

func (e userError) Error() string {
    return fmt.Sprintf("%v has a problem with their account", e.name)
}
```

It can then be used as an error:

```go
// Notice that this still takes the error interface as its return type
func sendSMS(msg, userName string) error {
    if !canSendUToUser(userName) {
        return userError{name: userName}
    }
}
```
