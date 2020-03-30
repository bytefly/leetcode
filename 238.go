package main

import "fmt"

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	l := make([]int, len(nums))
	r := make([]int, len(nums))

	l[0], r[len(r)-1] = 1, 1
	for i := 1; i < len(nums); i++ {
		l[i] = nums[i-1] * l[i-1]
	}
	for i := len(r) - 2; i >= 0; i-- {
		r[i] = nums[i+1] * r[i+1]
	}

	for i := 0; i < len(nums); i++ {
		ans[i] = l[i] * r[i]
	}
	return ans
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
