## The Initial Statement Of An If Block

An `if` conditional can have an "initial" statement. The variable(s) created in
the initial statement are <em>only</em> defined within the scope of the if body.

```go
if INITIAL_STATMENT; CONDITION {
}
```

\*\*Why Would I Use This?

This is just some syntactic sugar that Go offers to shorten up code in some
cases. For example, instead of writing:

```go
length := getLength(email)
if length < 1 {
    fmt.Println("Email is invalid")
}
```

We can do:

```go
if length := getLength(email); length < 1 {
    fmt.Println("Email is valid")
}
```

Not only is this code a bit shorter, but it also <em>removes length from the parent
scope</em>, which is convenient because we don't need it there - we only need access
to it while checking a condition.
