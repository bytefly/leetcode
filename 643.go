package main

import (
	"fmt"
)

func findMaxAverage(nums []int, k int) float64 {
	left, right := 0, k
	var sum int
	for i := left; i < right; i++ {
		sum += nums[i]
	}
	max := sum

	for right < len(nums) {
		sum -= nums[left]
		left++
		sum += nums[right]
		right++
		if sum > max {
			max = sum
		}
	}
	return float64(max) / float64(k)
}

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
}
