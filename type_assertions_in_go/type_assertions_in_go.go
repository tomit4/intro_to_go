package main

import (
	"fmt"
)

func getExpenseReport(e expense) (string, float64) {
	em, ok := e.(email)
	if ok {
		return em.toAddress(), em.cost()
	}
	s, ok := e.(sms)
	if ok {
		return s.toPhoneNumber(), s.cost()
	}
	return "", 0.0
}

// don't touch below this line

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * 0.05
	}
	return float64(len(e.body)) * 0.01
}

func (e email) toAddress() string {
	return fmt.Sprintf("Report: The email going to %s will cost %f64\n", e.body, e.cost())
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * 0.1
	}
	return float64(len(s.body)) * 0.03
}

func (s sms) toPhoneNumber() string {
	return fmt.Sprintf("Report: The sms going to %s will cost %f64\n", s.body, s.cost())
}

// func (i invalid) cost() float64 {
// return 0.0
// }

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
}

type sms struct {
	isSubscribed bool
	body         string
}

// main func not visible in video
