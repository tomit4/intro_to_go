package main

import "fmt"

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	// slice literal function, same as make function
	costsByDay := []float64{}
	for i := 0; i < len(costs); i++ {
		cost := costs[i]
		for cost.day >= len(costs) {
			costsByDay = append(costsByDay, 0.0)
		}
		costsByDay[cost.day] += cost.value

	}
	return costsByDay
}
