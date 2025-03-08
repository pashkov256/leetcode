package main

import "fmt"
func repeatedSubstringPattern(s string) bool {
	for i := 1; i <= len(s)/2; i++ {
		if len(s)%i == 0 {
			sub := s[:i]
			repeated := true
			for j := i; j < len(s); j += i {
				if s[j:j+i] != sub {
					repeated = false
					break
				}
			}
			if repeated {
				return true
			}
		}
	}
	return false
}


func main() {
	// fmt.Println(repeatedSubstringPattern("aabaaba"))
	// fmt.Println(repeatedSubstringPattern("abababab"))
	fmt.Println(repeatedSubstringPattern("abab"))
	
}