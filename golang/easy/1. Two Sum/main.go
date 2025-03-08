package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i,num := range nums {
		for j,subnum := range nums {
			if i == j {
				continue
			}
			if num + subnum == target {
				return []int{i,j}
			}
		}
	}
   return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2,7,11,15},9))
}