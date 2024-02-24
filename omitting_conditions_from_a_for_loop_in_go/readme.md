## Omitting Conditions From A For Loop In Go

Loops in Go can omit sections of a for loop. For example, the `CONDITION`
(middle part) can be omitted which causes the loop to run forever.

```go
for INITIAL;; AFTER {
    // do something forever
}
```
