package main

import (
	"fmt"
)

func rotate(nums []int, k int) {
	size := len(nums)
	k %= size

	for m := 0; m < k; m++ {
		t := nums[size-1]
		for i := size - 1; i > 0; i-- {
			nums[i] = nums[i-1]
		}
		nums[0] = t
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 1)
	fmt.Println(nums)

	nums = []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 2)
	fmt.Println(nums)

	nums = []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)

	nums = []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 4)
	fmt.Println(nums)

	nums = []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 5)
	fmt.Println(nums)

	nums = []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 6)
	fmt.Println(nums)

	nums = []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 1)
	fmt.Println(nums)
}
