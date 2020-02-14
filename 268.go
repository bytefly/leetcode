package main

import (
	"fmt"
)

func missingNumber(nums []int) int {
	var sum int

	for _, num := range nums {
		sum += num
	}
	return (len(nums)*(len(nums)+1))>>1 - sum
}

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 8, 1}))
}
