package main

import (
	"fmt"
	"sort"
)

func arrayRankTransform(arr []int) []int {
	ans := make([]int, len(arr))
	nums := make([]int, len(arr))
	m := make(map[int]int, len(arr))
	copy(nums, arr)

	sort.Ints(nums)
	next := 1
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; !ok {
			m[nums[i]] = next
			next++
		}
	}
	for i := 0; i < len(ans); i++ {
		ans[i] = m[arr[i]]
	}

	return ans
}

func main() {
	fmt.Println(arrayRankTransform([]int{40, 10, 20, 30}))
	fmt.Println(arrayRankTransform([]int{100, 100, 100}))
	fmt.Println(arrayRankTransform([]int{37, 12, 28, 9, 100, 56, 80, 5, 12}))
}
