package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	if len(flowerbed) == 1 && flowerbed[0] == 1{
		return false
	}

	plantedFlowersCount := n

	for i,num := range flowerbed {
		if num == 0 && i == 0{
			if i < len(flowerbed)-1 {		
				if flowerbed[i+1] == 0 {
					flowerbed[i] = 1
					plantedFlowersCount--
				}
			} 
			if len(flowerbed) == 1{
				plantedFlowersCount--
			}
		} else if i < len(flowerbed)-1 && i != 0{
			if flowerbed[i+1] == 0 && flowerbed[i-1] == 0 && num == 0{
				flowerbed[i] = 1
				plantedFlowersCount--
			}
		} else if i == len(flowerbed)-1 {
			if flowerbed[i-1] == 0 && num == 0 {
				flowerbed[i] = 1
				plantedFlowersCount--
			}
		} else if len(flowerbed) == 1 && num == 0 {		
			plantedFlowersCount--
		}
	}

	if plantedFlowersCount <= 0 {
		return true
	} else {
		return false
	}
}
func main() {
	fmt.Println(canPlaceFlowers([]int{0,0,0,0,1},2))
}