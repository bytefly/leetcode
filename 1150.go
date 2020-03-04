package main

import "fmt"

func isMajorityElement(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		if nums[left] == target && nums[right] == target {
			break
		}

		if nums[left] < target {
			left++
		} else if nums[right] > target {
			right--
		}
	}

	if left > right || right-left+1 <= len(nums)/2 {
		return false
	}

	return true
}

func main() {
	fmt.Println(isMajorityElement([]int{2, 4, 5, 5, 5, 5, 5, 6, 6}, 5))
	fmt.Println(isMajorityElement([]int{10, 100, 101, 101}, 101))
	fmt.Println(isMajorityElement([]int{10, 100, 101, 101}, 99))
	fmt.Println(isMajorityElement([]int{10, 100, 102, 102, 103}, 101))
	fmt.Println(isMajorityElement([]int{100}, 100))
	fmt.Println(isMajorityElement([]int{100}, 101))
}
