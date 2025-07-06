package main

import (
	"fmt"
)

func countConsistentStrings(allowed string, words []string) int {
	count := 0
	allowedMap := make(map[rune]struct{})

	for _, char := range allowed {
		allowedMap[char] = struct{}{}
	}

	for _, word := range words {
		consistent := true
		for _, char := range word {
			if _, exists := allowedMap[char]; !exists {
				consistent = false
				break
			}
		}
		if consistent {
			count++
		}

	}

	return count
}

func main() {
	var allowed = "cad"
	var words = []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}
	fmt.Println(countConsistentStrings(allowed, words))
}
