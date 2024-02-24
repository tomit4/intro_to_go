package main

import (
	"errors"
	"fmt"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string) ([]string, error) {
	allMessages := getMessageWithRetries()
	if plan == planPro {
		// Notice we are returning a slice instead of an array,
		// this is unlike in other languages, as we must be explicit
		// about the size of the slice
		return allMessages[:], nil
	}
	if plan == planFree {
		return allMessages[0:2], nil
	}
	return nil, errors.New("unsupported plan")
	// ?
}

func getMessageWithRetries() [3]string {
	return [3]string{
		"click here to sign up",
		"pretty please click here",
		"we beg you to sign up",
	}
}
