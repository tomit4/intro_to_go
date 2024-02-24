package main

import "fmt"

type divideError struct {
	dividend float64
}

func (de divideError) Error() string {
	return fmt.Sprintf("cannot divide %v by zero", de.dividend)
}
