package main

import (
	"fmt"
    "sort"
)

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	max := nums[len(nums)-3] * nums[len(nums)-2] * nums[len(nums)-1]
	num1 := nums[0] * nums[1] * nums[len(nums)-1]
	if max < num1 {
		max = num1
	}

	return max
}

func main() {
	fmt.Println(maximumProduct([]int{1, 2, 3}))
	fmt.Println(maximumProduct([]int{1, 2, 3, 4}))
	fmt.Println(maximumProduct([]int{-100, -99, 1, 1, 2, 3, 4}))
	fmt.Println(maximumProduct([]int{-100, -99, 0, 1, 2, 3, 4}))

	fmt.Println(maximumProduct([]int{-100, 0, 2, 3}))
}
