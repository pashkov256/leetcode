package main

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	var result []byte

	i,j := 0,0

	for i < len(word1) || j < len(word2) {
		if  i < len(word1)  {
			result = append(result, word1[i])
			i++
		}

		if j < len(word2) {
			result = append(result, word2[j])
			j++
		}
	}


	return string(result)
}

func main(){
	fmt.Println(mergeAlternately("abc","pqr"))
}