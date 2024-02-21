## Formatting Strings

**%v - Interpolate the default representation**

The %v variant prints the Go syntax representation of a value. You can usually
use this if you're unsure what else to use. That said, it's better to use the
type-specific variant if you can.

```go
fmt.Printf("I am %v years old", 10)
// I am 10 years old

fmt.Printf("I am %v years old", "way too many")
// I am way too many years old
```

**%s - Interpolate A String**

```go
fmt.Printf("I am %s years old", "way too many")
// I am way too many years old
```

**%d - Interpolate An Integer In Decimal Form**

```go
fmt.Printf("I am %d years old", 10)
// I am 10 years old
```

**%f - Interpolate A Decimal**

```go
fmt.Printf("I am %f years old", 10.523)
// I am 10.523000 years old
// The ".2" rounds the number to 2 decimal places
fmt.Printf("I am %.2f years old", 10.523)
// I am 10.53 years old
```
