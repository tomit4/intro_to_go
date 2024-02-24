package main

import "fmt"

func maxMessages(thresh float64) int {
	totalCost := 0.0
	for i := 0; ; i++ {
		totalCost += 1.0 + (0.01 * float64(i))
		// Stand in for initial omission of conditional
		if totalCost > thresh {
			return i
		}
	}
}
