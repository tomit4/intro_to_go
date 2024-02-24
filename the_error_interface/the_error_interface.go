package main

import "fmt"

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
	costForCustomer, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0.0, err
	}
	costForSpouse, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0.0, err
	}
	return costForCustomer + costForSpouse, nil
}

// don't edit below this line

func sendSMS(message string) (float64, error) {
	const maxTextLen = 25
	const costPerChar = .0002
	if len(message) > maxTextLen {
		return 0.0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * float64(len(message)), nil
}
