package main

import (
	"fmt"
)

func findErrorNums(nums []int) []int {
	var duplicate, lost int
	m := make([]bool, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		if !m[nums[i]] {
			m[nums[i]] = true
		} else {
			duplicate = nums[i]
		}
	}
	for i := 1; i < len(m); i++ {
		if m[i] == false {
			lost = i
			break
		}
	}
	return []int{duplicate, lost}
}

func main() {
	fmt.Println(findErrorNums([]int{1, 2, 2, 4}))
	fmt.Println(findErrorNums([]int{1, 1}))
}
