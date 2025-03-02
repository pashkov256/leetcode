package main

import "fmt"

func firstUniqChar(s string) int {
	lettersCountMap := make(map[rune]int)

	for _,letter := range s {
		lettersCountMap[letter] += 1
	}

	for i,letter := range s {
		if lettersCountMap[letter] == 1{
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(firstUniqChar("loveleetcode"))
}