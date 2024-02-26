## Concurrency

### What Is Concurrency?

Concurrency is the ability to perform multiple tasks at the same time.
Typically, our code is executed one line at a time, one after the other. This is
called <em>sequential execution</em> or <em>synchronous execution</em>.

```
            SYNCHRONOUS EXECUTION
-------------------->-------------------->
        TASK 1              TASK 2

            CONCURRENT EXECUTION
                   TASK 1
            ------------------->
            ------------------->
                   TASK 2
```

If the computer we're running our code on has multiple cores, we can even
execute multiple tasks at <em>exactly</em> the same time. If we're running on a
single core, a single code executes code at <em>almost</em> the same time by
switching between tasks very quickly. Either way, the code we write looks the
same in Go and takes advantage of whatever resources are availble.

### How Does Concurrency Work In Go?

Go was designed to be concurrent, which is a trait <em>fairly</em> unique to Go.
It excels at performing many tasks simultaneously safely using a simple syntax.

There isn't a popular programming language ine existence where spawning
concurrent execution is quite as elegant, at least in my opinion.

Concurrency is as simple as using the `go` keyword when calling a function:

```go
go doSomething()
```

In the example above, `doSomething()` will be executed concurrently with the
rest of the code in the function. The `go` keyword is used to spawn a new
[goroutine](https://golangdocs.com/goroutines-in-golang).
