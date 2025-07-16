package main

import "fmt"

//func twoSum(nums []int, target int) []int {
//	for i,num := range nums {
//		for j,subnum := range nums {
//			if i == j {
//				continue
//			}
//			if num + subnum == target {
//				return []int{i,j}
//			}
//		}
//	}
//   return []int{}
//}

// with two pointers
func twoSum(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1

	for left < right {
		summa := nums[left] + nums[right]

		if summa == target {
			return true
		}

		if summa > target {
			right--
		} else {
			left++
		}

	}

	return false
}

func main() {
	fmt.Println(twoSum([]int{1, 3, 4, 6, 8, 10, 13}, 13))
}
