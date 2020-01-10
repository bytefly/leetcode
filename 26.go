package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	size := 1
	if len(nums) == 0 {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[size-1] {
			nums[size] = nums[i]
			size++
		}
	}
	return size
}

func main() {
	//nums := []int{0, 0, 0, 0, 0}
	nums := []int{0, 0, 1, 1, 2, 2, 3, 3, 4}
	//nums := []int{1, 1, 2}
	//nums := []int{1, 2, 2}
	size := removeDuplicates(nums)

	for i := 0; i < size; i++ {
		fmt.Print(nums[i], " ")
	}
	fmt.Println()
}
