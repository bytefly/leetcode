package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	size := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[size] = nums[i]
			size++
		}
	}
	return size
}

func main() {
	nums := []int{0, 0, 0, 0, 0}
	//nums := []int{0, 0, 1, 1, 2, 2, 3, 3, 4}
	//nums := []int{1, 1, 2}
	//nums := []int{1, 2, 2}
	size := removeElement(nums, 0)

	for i := 0; i < size; i++ {
		fmt.Print(nums[i], " ")
	}
	fmt.Println()
}
