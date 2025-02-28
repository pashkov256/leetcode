package main

import "fmt"


func countSymbolInWord(word string)map[rune]int{
	countMap := make(map[rune]int,0)

	for _,s := range word {
		countMap[s] += 1
	}

	return countMap
}


func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := countSymbolInWord(s)
	tMap := countSymbolInWord(t)


	for key,value := range sMap {
		if tMap[key] != value {
			return false
		} 
	}

	return true
}

func main() {
	fmt.Println(isAnagram("anagram","nagaram"))//true
}