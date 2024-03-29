package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func getCounts(userIDs []string) map[string]int {
	counts := make(map[string]int)
	for _, userID := range userIDs {
		count := counts[userID]
		count++
		counts[userID] = count
	}
	return counts
}
