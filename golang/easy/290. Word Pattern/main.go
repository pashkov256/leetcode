package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	wordsSlice := strings.Split(s," ")
	if len(pattern) != len(wordsSlice) {
		return false
	}
	patternMap := make(map[string]string)
	usedWords := make(map[string]struct{})

	for i,word := range wordsSlice {
		patternLetter := string(pattern[i])
		if patternValue,exists := patternMap[patternLetter];!exists{
			if _,exists := usedWords[word];!exists {
				patternMap[patternLetter] = word 
				usedWords[word] = struct{}{}
			} else {
				return false 
			}
		} else {
			if word != patternValue {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Println(wordPattern("abba","dog cat cat dog"))
	fmt.Println(wordPattern("abba","dog dog dog dog"))
}