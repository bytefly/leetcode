package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	ans := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			ans = i
			break
		}
		if target < nums[i] {
			ans = i
			break
		}
	}

	if ans < 0 {
		ans = len(nums)
	}
	return ans
}

func main() {
	fmt.Println(searchInsert([]int{}, 1))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
}
