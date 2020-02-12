package main

import (
	"fmt"
)

func findUnsortedSubarray(nums []int) int {
	var left, right int
	min := int(^uint32(0) >> 1)
	max := int(^min)

	// find the minimum decreae value from start
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] && nums[i] < min {
			min = nums[i]
		}
	}
	// find the maximum increase value from end
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i+1] < nums[i] && nums[i] > max {
			max = nums[i]
		}
	}
	// find the right place of minimum value
	for i := 0; i < len(nums); i++ {
		if nums[i] > min {
			left = i
			break
		}
	}
	// find the right place of maximum value
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < max {
			right = i
			break
		}
	}
	// no moving for ordered array
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
	fmt.Println(findUnsortedSubarray([]int{1, 3, 8, 2, 6, 3, 5, 9}))
}
