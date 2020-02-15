package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	cnt := 1
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		if cnt == 0 {
			ans = nums[i]
			cnt = 1
			continue
		}
		if ans != nums[i] {
			cnt--
		} else {
			cnt++
		}
	}
	return ans
}

func main() {
	fmt.Println(majorityElement([]int{3}))
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{10, 9, 9, 9, 10}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
