package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	ans := make([]int, 0)
	i, j := 0, 0
	for {
		if i == m || j == n {
			break
		}
		if nums1[i] <= nums2[j] {
			ans = append(ans, nums1[i])
			i++
		} else {
			ans = append(ans, nums2[j])
			j++
		}
	}

	if i < m {
		ans = append(ans, nums1[i:]...)
	}
	if j < n {
		ans = append(ans, nums2[j:]...)
	}
	copy(nums1, ans)
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}
