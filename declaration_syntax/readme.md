## Declaration Syntax

Developers often wonder why the declaration syntax in Go is different from the
tradition established in the C family of languages.

**C-Style Syntax**

The C language describes types with an expression including the name to be
declared, and states what type that expression will have.

```c
int y;
```

The code above declares `y` as an `int`. In general, the type goes on the left
and the expression on the right.

Interestingly, teh creators of the Go language agreed that the C-style of
declaring types ins ignatures gets confusing really fast - take a look at this
nightmare.

```c
int (*fp)(int (*ff)(int x, int y), int b)
```

**Go-Style Syntax**

Go's declarations are clear, you just read them left to right, just like you
would in English.

```go
x int
p *int
a [3]int
```

It's nice for more complex signatures, it makes them easier to read.

```go
/* Here we have a callback that itself takes two integers as arguments, returns
an int type, and the second argument passed to the outer function is also an
int, with the outer function's return type being an int */
f func(func(int, int) int, int) int
```
