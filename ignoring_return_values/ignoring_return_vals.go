package main

import "fmt"

func main() {
	// firstName, lastName := getNames()
	// Ignoring the lastName variable,
	// the compiler actually removes it since it's never used
	firstName, _ := getNames()
	fmt.Println("Welcome to Textio", firstName)
}

// don't edit below this line

func getNames() (string, string) {
	return "John", "Doe"
}
