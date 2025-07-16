package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	normalizeStr := strings.ToLower(s)

	for left < right {
		if !isAlphanumeric(normalizeStr[left]) {
			left++
			continue
		}
		if !isAlphanumeric(normalizeStr[right]) {
			right--
			continue
		}

		if normalizeStr[left] == normalizeStr[right] {
			left++
			right--
		} else {
			return false
		}
	}

	return true
}

func isAlphanumeric(l byte) bool {
	return (l >= 'a' && l <= 'z') || (l >= '0' && l <= '9')
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) //amanaplanacanalpanama
}
