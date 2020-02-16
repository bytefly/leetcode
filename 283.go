package main

import (
	"fmt"
)

func moveZeroes(nums []int) {
	last := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if last >= 0 {
				nums[last], nums[i] = nums[i], nums[last]
				last += 1
			}
		} else {
			if last < 0 {
				last = i
			}
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)

	nums = []int{1, 0, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)

	nums = []int{1, 0, 0, 3, 0}
	moveZeroes(nums)
	fmt.Println(nums)

	nums = []int{0, 1, 1, 2, 3}
	moveZeroes(nums)
	fmt.Println(nums)
}
