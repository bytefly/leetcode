package main

import (
	"fmt"
)

func intersection(nums1 []int, nums2 []int) []int {
	var ans []int
	m := make(map[int]bool)
	n := make(map[int]bool)

	for _, v := range nums1 {
		m[v] = true
	}
	for _, v := range nums2 {
		if _, ok := m[v]; ok {
			n[v] = true
		}
	}
	for k, _ := range n {
		ans = append(ans, k)
	}
	return ans
}

func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
