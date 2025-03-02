package main

import "fmt"

func findTheDifference(s string, t string) byte {
	letterSMap := make(map[byte]struct{})
	letterTMap := make(map[byte]struct{})


	for _,letter := range s {
		if _,ok := letterSMap[byte(letter)]; !ok {
			letterSMap[byte(letter)] = struct{}{}
		}
	}

	for _,letter := range t {
		if _,ok := letterTMap[byte(letter)]; !ok {
			letterTMap[byte(letter)] = struct{}{}
		}
	}


	fmt.Println(len(letterSMap))

	if len(letterTMap) > len(letterSMap){
		for key := range letterTMap {
			if _,ok :=  letterSMap[key];!ok{
				fmt.Println(string(key))
				return key
			} 
		}
	} else {
		for key := range letterSMap {
			if _,ok :=  letterTMap[key];!ok{
				fmt.Println(string(key))
				return key
			} 
		}
	}

	


	return 1
}

func main() {
	fmt.Println(findTheDifference("a","aa"))
}