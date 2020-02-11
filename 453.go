package main

import (
	"fmt"
)

func minMoves(nums []int) int {
	var ans int

	min := nums[0]
	for _, i := range nums {
		if i < min {
			min = i
		}
	}
	for _, i := range nums {
		ans += i - min
	}

	return ans
}

func main() {
	fmt.Println(minMoves([]int{1, 3}))
	fmt.Println(minMoves([]int{1, 2, 3}))
	fmt.Println(minMoves([]int{1, 2, 1}))
	fmt.Println(minMoves([]int{1, 2, 2}))
	fmt.Println(minMoves([]int{1, 4, 4}))
	fmt.Println(minMoves([]int{1, 2, 3, 4, 5, 6}))
}
