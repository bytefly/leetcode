package main

import (
	"fmt"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	ints := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := ints[nums[i]]; ok {
			return true
		} else {
			ints[nums[i]]++
		}
		if len(ints) > k {
			delete(ints, nums[i-k])
		}
	}
	return false
}

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
	fmt.Println(containsNearbyDuplicate([]int{1}, 0))
	fmt.Println(containsNearbyDuplicate([]int{99, 99}, 2))
}
