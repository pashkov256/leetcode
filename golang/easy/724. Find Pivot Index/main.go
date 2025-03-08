package main

import "fmt"

func pivotIndex(nums []int) int {

	for i,num := range nums {
		leftSum := 0
		rightSum := 0
		isNotZero := false
		fmt.Println("num",num)
		//i != 0 && i != len(nums)-1
	 	if len(nums) >= 3{
			for _,rightNum := range nums[i+1:] {
				if rightNum != 0 {
					isNotZero = true
				}
				rightSum += rightNum
			}
			for _,lefttNum := range nums[:i] {
				if lefttNum != 0 {
					isNotZero = true
				}
				leftSum += lefttNum
			}
		}
		
		fmt.Println("leftSum",leftSum,"\n","rightSum",rightSum)
		fmt.Println("isNotZero",isNotZero)
		fmt.Println()

		if leftSum == rightSum{
			if isNotZero {
				return i
			} else if leftSum + rightSum != 0 {
						return i
			}
	
		}
		// if leftSum == rightSum &&  leftSum + rightSum != 0{
		// 	return i
		// }
	}

	return -1
}

func main() {
	// fmt.Println(pivotIndex([]int{1,7,3,6,5,6}))
	// fmt.Println(pivotIndex([]int{2,1,-1}))
	fmt.Println(pivotIndex([]int{-1,0,0,0,0,0}))
}