package main

import (
	"fmt"
)

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		beginSearch := false
		ans[i] = -1
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				beginSearch = true
				continue
			}
			if nums1[i] < nums2[j] && beginSearch {
				ans[i] = nums2[j]
				break
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Println(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
}
