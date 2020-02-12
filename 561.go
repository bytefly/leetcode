package main

import (
	"fmt"
	"sort"
)

func arrayPairSum(nums []int) int {
	var ans int
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		ans += nums[i]
		fmt.Println(nums)
	}
	return ans
}

func main() {
	fmt.Println(arrayPairSum([]int{1, 4, 3, 2}))
	fmt.Println(arrayPairSum([]int{5, 6, 7, 8, 9, 10, 11, 12}))
}
