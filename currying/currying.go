package main

import (
	"errors"
	"fmt"
)

// getLogger takes a function that formats two strings into
// a single string and returns a function that formats two strings
// the result instead of returning it

func getLogger(formatter func(string, string) string) func(string, string) {
	return func(a string, b string) {
		fmt.Println(formatter(a, b))
	}
}
