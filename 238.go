package main

import "fmt"

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))

	ans[0] = 1
	for i := 1; i < len(nums); i++ {
		ans[i] = nums[i-1] * ans[i-1]
	}

	r := 1
	for i := len(nums) - 2; i >= 0; i-- {
		r *= nums[i+1]
		ans[i] *= r
	}

	return ans
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
