package main

import "fmt"

func main() {
	// declare here
	// These two are equivlanet declarations of an empty string:
	// this is considered very verbose, as it allows you to specify the type
	// var empty string
	// Usually you will declare variables like so which infers the type:
	// empty := ""
	congrats := "happy birthday"
	fmt.Println(congrats)
}
