package main

import "fmt"

func main() {
	// NOTE: constant declarations  cannot use the short form :=.
	const premiumPlanName = "Premium Plan"
	// errors here, as constants cannot be redeclared
	// premiumPlanName = "Basic Plan"
	const basicPlanName = "Basic Plan"

	// don't edit below this line

	fmt.Println("plan:", premiumPlanName)
	fmt.Println("plan:", basicPlanName)
}
