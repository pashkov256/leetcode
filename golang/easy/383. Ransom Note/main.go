package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
  }
	magazineLettersMap := make(map[rune]int)

	for _,letter := range magazine {
		magazineLettersMap[letter] += 1
	}


	for _,letter := range ransomNote {
		if valueMap,exists := magazineLettersMap[letter];exists {
			if valueMap <= 0 {
				return false
			}
			magazineLettersMap[letter] -= 1
		}else {
			return false
		}
	}

	return true
}

func main() {
	// fmt.Println(canConstruct("aa","aab"))
	fmt.Println(canConstruct("a","b"))
}