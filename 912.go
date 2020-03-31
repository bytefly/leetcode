package main

import (
	"fmt"
	"sort"
)

func sortArray(nums []int) []int {
	ans := make([]int, len(nums))
	copy(ans, nums)
	sort.Ints(ans)
	return ans
}

func main() {
	fmt.Println(sortArray([]int{5, 2, 3, 1}))
	fmt.Println(sortArray([]int{5, 1, 1, 2, 0, 0}))
}
