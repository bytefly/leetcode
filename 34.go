package main

import "fmt"

func searchLeft(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if left < len(nums) && nums[left] == target {
		return left
	}
	return -1
}

func searchRight(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if left == 0 && nums[0] != target {
		return -1
	}
	if nums[left-1] == target {
		return left - 1
	}

	return -1
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	ans := make([]int, 2)
	ans[0], ans[1] = searchLeft(nums, target), searchRight(nums, target)
	return ans
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 8}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 8}, 5))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 8}, 7))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 4))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 12))
}
