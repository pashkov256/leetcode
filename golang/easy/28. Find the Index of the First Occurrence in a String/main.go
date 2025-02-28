package main

import "fmt"

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	nLen := len(needle)

	for i := range haystack {
		if (i+nLen) <= len(haystack){
			if haystack[i:i+nLen] == needle {
				return i
			}
		}
	}

	return -1
}

func main() {
	fmt.Println(strStr("leetcode","leeto"))
}