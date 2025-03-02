package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	nMap := make(map[int]struct{})
	var intersectionSlice []int 


	for _,num := range nums1 {
		nMap[num] = struct{}{}
	}

	for _,num := range nums2 {
		if _,exists := nMap[num];exists{
			delete(nMap,num)
			intersectionSlice = append(intersectionSlice, num)
		}
	}


	return intersectionSlice
}



//РЕШЕНИЕ 1
// func containts(slice []int,item int) bool {
// 	for _,elem := range slice {
// 		if elem == item {
// 			return true
// 		} 
// 	}
// 	return false
// }

// func intersection(nums1 []int, nums2 []int) []int {
// 	intersectionSlice := make([]int,0,8)

// 	for _,num := range nums1 {
// 		if !containts(intersectionSlice,num) {
// 			if containts(nums2,num) {
// 				intersectionSlice = append(intersectionSlice, num)
// 			}
// 		}
// 	}

// 	return intersectionSlice
// }

func main() {
	fmt.Println(intersection([]int{1,2,2,1},[]int{2,2}))
}