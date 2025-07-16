package main

import "fmt"

func maxArea(height []int) int {
	maxAreaInt := 0

	left, right := 0, len(height)-1

	for left < right {
		minHeight := height[right]
		lengthDiff := right - left

		if height[left] < height[right] {
			minHeight = height[left]
			left++
		} else {
			right--
		}

		areaWater := minHeight * lengthDiff

		if maxAreaInt < areaWater {
			maxAreaInt = areaWater
		}

	}

	return maxAreaInt
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{4, 3, 2, 1, 4})) //16
}
