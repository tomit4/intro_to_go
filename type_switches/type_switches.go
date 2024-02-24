package main

import (
	"fmt"
)

func getExpenseReport(e expense) (string, float64) {
    switch v := e.(type) {
    case email:
	return v.toAddress, v.cost()
    case sms:
	return v.toPhoneNumber, v.cost()
    default:
	return "", 0.0
}

// don't touch below this line
/*
func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * 0.05
	}
	return float64(len(e.body)) * 0.01
}
*/
