## Type Sizes

Ints, uints, floats, and complex numbers all have type sizes.

int int8 int16 int32 int64 // whole numbers

uint uint8 uint16 uint32 uint64 // positive whole numbers

float32 float64 // decimal point

complex64 complex128 // imaginary numbers (rare)

The size (8, 16, 32, 64, 128, etc) indicates how many bits in memory will be
used to store the variable. The default `int` and `uint` types are just aliases
that refer to their respective 32 or 64 bit sizes depending on the environment
of the user.

The standard sizes that should be used unless the developer has a specific need
are:

- int
- uint
- float64
- complex128 (again rare)

Some types can be converted the following way:

```go
tempeartureInt := 88
temperatureFloat := float64(temperatureInt)
```
