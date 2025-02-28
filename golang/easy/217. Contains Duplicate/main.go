package main

import "fmt"

func containsDuplicate(nums []int) bool {
	existing := make(map[int]struct{},0)

	for _,num := range nums{
		if _,ok := existing[num];ok {
			return true
		}
		existing[num] = struct{}{}
		
	}

	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1,1,1,3,3,4,3,2,4,2}))
}