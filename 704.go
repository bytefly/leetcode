package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{1}, 1))
	fmt.Println(search([]int{1}, 0))
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 0))
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, -1))
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 12))
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, -2))
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 13))
}
