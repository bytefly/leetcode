package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	ans := nums[0]

	for i := 1; i < len(nums); i++ {
		ans ^= nums[i]
	}
	return ans
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2, 4, 0}))
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2, 4, -2}))
}
