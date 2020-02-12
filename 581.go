package main

import (
	"fmt"
	"sort"
)

func findUnsortedSubarray(nums []int) int {
	newNums := make([]int, len(nums))
	copy(newNums, nums)
	sort.Ints(newNums)

	left, right := 0, len(nums)-1
	for left < right {
		if newNums[left] == nums[left] {
			left++
		}
		if newNums[right] == nums[right] {
			right--
		}

		if newNums[left] != nums[left] && newNums[right] != nums[right] {
			break
		}
	}

	if left == right {
		return 0
	}

	return right - left + 1
}

func main() {
	fmt.Println(findUnsortedSubarray([]int{6, 6, 2, 8, 10, 15}))
	fmt.Println(findUnsortedSubarray([]int{2, 6, 7, 8, 10, 15, 10}))
	fmt.Println(findUnsortedSubarray([]int{2, 6, 7, 8, 15, 10, 10}))
	fmt.Println(findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
	fmt.Println(findUnsortedSubarray([]int{2, 6, 8, 10, 15}))
	fmt.Println(findUnsortedSubarray([]int{2, 6, 10, 8, 15}))
	fmt.Println(findUnsortedSubarray([]int{6, 2, 8, 10, 15}))
	fmt.Println(findUnsortedSubarray([]int{2, 6, 8, 10, 15, 14}))
	fmt.Println(findUnsortedSubarray([]int{6, 2, 8, 10, 15, 14}))
}
