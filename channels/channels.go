package main

import (
	"fmt"
	"time"
)

func filterOldEmails(emails []email) {
	isOldChan := make(chan bool)

	go func() {
		/*
		 *Originally outside of the goroutine, this for loop
		 * never exits because the channel is waiting for isOldChan
		 * to be assigned a value (channel is blocking)
		 */
		for _, e := range emails {
			if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
				isOldChan <- true
				continue
			} else {
				isOldChan <- false
			}
		}
	}()

	isOld := <-isOldChan
	fmt.Println("email 1 is old", isOld)
	isOld = <-isOldChan
	fmt.Println("email 2 is old", isOld)
	isOld = <-isOldChan
	fmt.Println("email 3 is old", isOld)
}
