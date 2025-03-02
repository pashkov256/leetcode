package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1


	for left <= right {
		mid := left + (right-left) / 2

		if nums[mid] == target {
			return nums[mid]
		} else if nums[mid] < target { // Если элемент меньше цели, ищем справа
			left = mid + 1
	  } else { // Если элемент больше цели, ищем слева
			right = mid - 1
	  }

	
	}
	return 1
}

func main() {
	sortedArray := []int{1, 3, 5, 6}
	target := 2

	index := searchInsert(sortedArray, target)
	fmt.Println(index)
}