package main

import (
	"fmt"
)

func intersect(nums1 []int, nums2 []int) []int {
	var ans []int
	m := make(map[int]int)

	for _, v := range nums1 {
		m[v]++
	}
	for _, v := range nums2 {
		if _, ok := m[v]; ok {
			ans = append(ans, v)
			m[v]--
			if m[v] == 0 {
				delete(m, v)
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
