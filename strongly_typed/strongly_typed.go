package main

import "fmt"

func main() {
	var username string = "wagslane"
	// var password int = 20947382822

	// don't edit below this line
	// throws error cuz you can't put two types together
	// fmt.Println("Authorization: Basic", username+":"+password)
	// Thusly change it to a string and you'll be good to go
	var password string = "20947382822"
	fmt.Println("Authorization: Basic", username+":"+password)
}
