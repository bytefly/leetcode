package main

import "fmt"

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	t := right

	if len(nums) == 0 {
		return -1
	}
	//find the largest position
	if nums[t] < nums[left] {
		for left <= t {
			mid := left + (t-left)>>1
			if nums[mid] > nums[mid+1] {
				t = mid
				break
			} else if nums[mid] > nums[right] {
				left = mid + 1
			} else {
				t = mid - 1
			}
		}
	}

	if target >= nums[0] && target <= nums[t] {
		left, right = 0, t
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
	fmt.Println(search([]int{3, 4, 5, 6, 7, 1, 2}, 4))
}
