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

func groupAnagrams(strs []string) [][]string {
	var res = make([][]string,0,len(strs)/2)
	usedAnagram := make(map[string]struct{})
	for i,str := range strs {
		if _,exists := usedAnagram[str];exists{
			continue
		}
		resUnd := []string{str}
		for _,otherStr := range strs[i+1:] {
			if isAnagram(str,otherStr) {
				resUnd = append(resUnd, otherStr)
				usedAnagram[otherStr] = struct{}{}
			}
		}
		res = append(res, resUnd)
	}
	
	return res
}

func main() {
	// fmt.Println(groupAnagrams([]string{"eat","tea","tan","ate","nat","bat"}))
	fmt.Printf("%#v",groupAnagrams([]string{""}))
}