package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		sum := numbers[left] + numbers[right]

		if sum == target {
			return []int{left + 1, right + 1}
		}

		if sum < target {
			left++
		} else {
			right--
		}
	}

	return []int{-1, -1}
}

func main() {
	fmt.Println(twoSum([]int{2, 3, 4}, 6))
}
