package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	wordDictMap := make(map[string]struct{})
	for range s {
		for _, word := range wordDict {
			if len(s) >= len(word) {
				if s[:len(word)] == word {
					wordDictMap[word] = struct{}{}
					s = s[len(word):]
				}
			}

		}
	}

	if len(wordDictMap) == len(wordDict) {
		return true
	}

	return false
}

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	// fmt.Println(wordBreak("bb",[]string{"a","b","bbb","bbbb"}))
}