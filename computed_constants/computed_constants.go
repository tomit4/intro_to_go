package main

import "fmt"

func main() {
	const secondsInMinute = 60
	const minutesInHour = 60
	// See here that we have an undefined constant
	// const secondsInHour = // ?
	// This also won't work, unlike in JS, you cannot declare an undefined constant
	// const secondsInHour
	// NOTE that constants are not computed at run time, but rather at compile time
	const secondsInHour = secondsInMinute * minutesInHour

	// don't edit below this line
	fmt.Println("number of seconds in an hour:", secondsInHour)

}
