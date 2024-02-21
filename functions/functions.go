package main

import "fmt"

// Throws an error as we have not defined the types in our arguments declaration
/*
func concat(s1, s2) string {
	return s1 + s2
}
*/

// This will work
func concat(s1 string, s2 string) string {
	return s1 + s2
}

// Multiple Parameters
func add(x, y int) int {
	return x + y
}
func createUser(firstName, lastName string, age int) {
	fmt.Printf("My name is %s %s, and I am %d years old\n", firstName, lastName, age)
}

// don't touch below this line

func main() {
	createUser("Lane", "Wagner", 27)
	fmt.Println(concat("Lane,", " happy birthday!"))
	fmt.Println(concat("Elon,", " hope that Tesla thing works out"))
	fmt.Println(concat("Go", " is fantastic"))
}
