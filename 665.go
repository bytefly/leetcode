package main

import (
	"fmt"
)

func checkPossibility(nums []int) bool {
	var count int
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			if i-2 >= 0 && nums[i-2] > nums[i] {
				nums[i] = nums[i-1]
			} else {
				nums[i-1] = nums[i]
			}
			count++
		}
	}
	return count <= 1
}

func main() {
	fmt.Println(checkPossibility([]int{4, 2}))
	fmt.Println(checkPossibility([]int{4, 2, 3}))
	fmt.Println(checkPossibility([]int{4, 2, 1}))
	fmt.Println(checkPossibility([]int{3, 4, 2, 3}))
	fmt.Println(checkPossibility([]int{3, 3, 2, 2}))
	fmt.Println(checkPossibility([]int{-1, 4, 2, 3}))
	fmt.Println(checkPossibility([]int{4, 2, 3, 2, 4, 3}))
}
