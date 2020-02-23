package main

import (
	"fmt"
)

func decompressRLElist(nums []int) []int {
	var ans []int
	for i := 0; i < len(nums); i += 2 {
		for j := 0; j < nums[i]; j++ {
			ans = append(ans, nums[i+1])
		}
	}
	return ans
}

func main() {
	fmt.Println(decompressRLElist([]int{1, 2, 3, 4}))
}
