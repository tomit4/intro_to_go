package main

import (
	"fmt"
	"time"
)

func sendEmail(message string) {
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	/* The go keyword above ensures that the
	 * above function occurs in a new thread, a "goroutine",
	 * ensuring that the function is not blocking, and therefore that the
	 * next line will be executed on its own thread (i.e. sent happens before received)
	 */
	fmt.Printf("Email sent: '%s'\n", message)
}
