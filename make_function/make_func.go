package main

import "fmt"

func getMessageCosts(messages []string) []float64 {
	// let costs = []
	// costs.length = messages.length
	costs := make([]float64, len(messages))
	// messages.forEach(message => ...)
	for i := 0; i < len(messages); i++ {
		message := messages[i]
		// const cost = message.length * 0.01
		cost := float64(len(message)) * 0.01
		costs[i] = cost
	}
	return costs
}
