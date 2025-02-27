package main

import "fmt"

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	nLen := len(needle)

	for i := range haystack {

		fmt.Println("LENG",len(haystack))
		fmt.Println("MAXEND",i+nLen)
		if (i+nLen) <= len(haystack){
			fmt.Println(haystack[i:i+nLen])
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