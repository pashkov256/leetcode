package main

import "fmt"



func majorityElement(nums []int) int {
	candidate := nums[0]
	count := 1

	for i:=1;i< len(nums);i++{		
		if count == 0{
			candidate = nums[i]
		}

		if nums[i] == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}

func main() {
	fmt.Println(majorityElement([]int{2,2,1,1,1,2,2}))
}