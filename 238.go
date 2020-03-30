package main

import "fmt"

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	for i := 0; i < len(ans); i++ {
		ans[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			ans[j] *= nums[i]
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			ans[j] *= nums[i]
		}
	}

	return ans
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
