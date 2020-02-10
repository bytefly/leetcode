package main

import (
	"fmt"
)

func findDisappearedNumbers(nums []int) []int {
	for _, j := range nums {
		if j < 0 {
			j = -j
		}
		if nums[j-1] > 0 {
			nums[j-1] = -nums[j-1]
		}
	}

	var ans []int
	for i, j := range nums {
		if j > 0 {
			ans = append(ans, i+1)
		}
	}
	return ans
}

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 1, 5, 1}))
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 1, 5, 6}))
}
