package main

import (
	"fmt"
	"sort"
)

func arrayRankTransform(arr []int) []int {
	ans := make([]int, len(arr))
	nums := make([]int, len(arr))
	copy(nums, arr)

	sort.Ints(nums)
	for i := 0; i < len(ans); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[j] < arr[i] && (j == 0 || nums[j] != nums[j-1]) {
				ans[i]++
			}
		}
		ans[i]++
	}

	return ans
}

func main() {
	fmt.Println(arrayRankTransform([]int{40, 10, 20, 30}))
	fmt.Println(arrayRankTransform([]int{100, 100, 100}))
	fmt.Println(arrayRankTransform([]int{37, 12, 28, 9, 100, 56, 80, 5, 12}))
}
