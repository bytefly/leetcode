package main

import "fmt"

func massage(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	ans := make([]int, len(nums))
	ans[0], ans[1] = nums[0], nums[1]
	if ans[1] < ans[0] {
		ans[1] = ans[0]
	}

	for i := 2; i < len(nums); i++ {
		ans[i] = ans[i-2] + nums[i]
		if ans[i] < ans[i-1] {
			ans[i] = ans[i-1]
		}
	}

	return ans[len(nums)-1]
}

func main() {
	fmt.Println(massage([]int{1, 2, 3, 1}))
	fmt.Println(massage([]int{2, 7, 9, 3, 1}))
	fmt.Println(massage([]int{2, 1, 4, 5, 3, 1, 1, 3}))
}
