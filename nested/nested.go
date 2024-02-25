package main

import (
	"fmt"
)

// rune stands in for a single character in go
func getNameCounts(names []string) map[rune]map[string]int {
	counts := make(map[rune]map[string]int)
	for _, name := range names {
		if name == "" {
			continue
		}
		firstChar := rune(name[0])
		_, ok := counts[firstChar]
		// If the key by character doesn't exist...
		if !ok {
			// initialize it
			counts[firstChar] = make(map[string]int)
		}
		// count it
		counts[firstChar][name]++
	}
	return counts
}
