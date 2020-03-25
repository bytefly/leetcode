package main

import "fmt"

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	t := right

	if len(nums) == 0 {
		return -1
	}
	if nums[t] < nums[left] {
		for t >= 0 {
			if nums[t] >= nums[0] {
				break
			}
			t--
		}
	}

	if target >= nums[0] && target <= nums[t] {
		right = t
	} else if t < len(nums)-1 && target >= nums[t+1] && target <= nums[right] {
		left, right = t+1, len(nums)-1
	} else {
		return -1
	}

	for left <= right {
		mid := left + (right-left)>>1
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] > target:
			right = mid - 1
		case nums[mid] < target:
			left = mid + 1
		}
	}

	return -1
}

func main() {
	fmt.Println(search([]int{3, 1}, 1))
	fmt.Println(search([]int{0}, 0))
	fmt.Println(search([]int{0}, 1))
	fmt.Println(search([]int{}, 1))
	fmt.Println(search([]int{0, 1, 2, 4, 5, 6}, 0))
	fmt.Println(search([]int{0, 1, 2, 4, 5, 6}, 6))
	fmt.Println(search([]int{0, 1, 2, 4, 5, 6}, 5))
	fmt.Println(search([]int{4, 5, 6, 0, 1, 2}, 0))
	fmt.Println(search([]int{4, 5, 6, 0, 1, 2}, 1))
	fmt.Println(search([]int{4, 5, 6, 0, 1, 2}, 5))
	fmt.Println(search([]int{4, 5, 6, 0, 1, 2}, 3))
}
