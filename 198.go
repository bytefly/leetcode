package main

import (
	"fmt"
	"math"
)

func rob(nums []int) int {
	var prevMax, max float64
	for i := 0; i < len(nums); i++ {
		t := max
		max = math.Max(prevMax+float64(nums[i]), max)
		prevMax = t
	}
	return int(max)
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
	fmt.Println(rob([]int{2, 1, 1, 2}))
	fmt.Println(rob([]int{1}))
	fmt.Println(rob([]int{1, 3, 1, 3, 100}))
}
