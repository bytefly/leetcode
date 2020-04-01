package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	var ans int

	cnt, n := 0, len(nums)
	for i := n/2 - 1; i >= 0; i-- {
		heaplize(nums, i, n)
	}

	for i := n - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		cnt++
		if cnt == k {
			ans = nums[i]
			break
		}
		heaplize(nums, 0, i)
	}

	return ans
}

func heaplize(nums []int, i, end int) {
	left, right := 2*i+1, 2*i+2
	large := i
	if left < end && nums[left] > nums[i] {
		large = left
	}
	if right < end && nums[right] > nums[large] {
		large = right
	}

	if large != i {
		nums[i], nums[large] = nums[large], nums[i]
		heaplize(nums, large, end)
	}
}

func main() {
	fmt.Println(findKthLargest([]int{3}, 1))
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
