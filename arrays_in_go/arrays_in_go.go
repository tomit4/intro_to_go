package main

import "fmt"

const (
	retry1 = "click here to sign up"
	retry2 = "pretty please click here"
	retry3 = "we beg you to sign up"
)

func getMessageWithRetries() [3]string {
	return [3]string{
		"click here to sign up",
		"pretty please click here",
		"we beg you to sign up",
	}
}

func main() {
	message := getMessageWithRetries()
	fmt.Println(message)
}
