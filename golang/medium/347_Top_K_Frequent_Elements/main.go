package main

import (
	"fmt"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	frequentMap := make(map[int]int)
	for _, num := range nums {
		frequentMap[num]++
	}

	pairs := make([][]int, 0, len(frequentMap))
	for key, value := range frequentMap {
		pairs = append(pairs, []int{key, value})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] > pairs[j][1]
	})

	result := make([]int, 0, k)
	for i := 0; i < k; i++ {
		result = append(result, pairs[i][0])
	}

	return result
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}
