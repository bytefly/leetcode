package main

import (
	"fmt"
)

func findDisappearedNumbers(nums []int) []int {
	numMap := make([]int, len(nums))
	for _, v := range nums {
		numMap[v-1] = 1
	}

	var ans []int
	for k, v := range numMap {
		if v == 0 {
			ans = append(ans, k+1)
		}
	}
	return ans
}

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 1, 5, 1}))
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 1, 5, 6}))
}
