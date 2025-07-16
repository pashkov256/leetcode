package main

import (
	"fmt"
	"strings"
)

func reverseVowels(s string) string {
	vowels := map[string]struct{}{
		"a": {},
		"e": {},
		"i": {},
		"o": {},
		"u": {},
	}
	left, right := 0, len(s)-1

	reversedVowelsStr := []rune(s)

	for left < right {
		if _, leftCharIsVowel := vowels[strings.ToLower(string(s[left]))]; !leftCharIsVowel {
			left++
			continue
		}
		if _, rightCharIsVowel := vowels[strings.ToLower(string(s[right]))]; !rightCharIsVowel {
			right--
			continue
		}

		reversedVowelsStr[left], reversedVowelsStr[right] = rune(s[right]), rune(s[left])

		left++
		right--
	}

	return string(reversedVowelsStr)
}

func main() {
	fmt.Println(reverseVowels("IceCreAm")) //AceCreIm
}
