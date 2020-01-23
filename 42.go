package main

import (
	"fmt"
)

func trap(height []int) int {
	var maxHeight, ans int

	//find the max height
	for i := 0; i < len(height); i++ {
		if height[i] > maxHeight {
			maxHeight = height[i]
		}
	}
	if maxHeight < 1 || len(height) <= 1 {
		return 0
	}

	for i := 1; i <= maxHeight; i++ {
		start, end := -1, -1
		for j := 1; j < len(height); j++ {
			if height[j-1] > height[j] && height[j] < i && height[j-1] >= i {
				start = j - 1
			}
			if height[j-1] <= height[j] && height[j] >= i {
				end = j
				if start >= 0 {
					ans += (end - 1 - start)
					fmt.Println(start, end, i)
				}
				start, end = -1, -1
			}
		}
	}
	return ans
}

func main() {
	keys := [][]int{
		[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
		[]int{2, 0, 2},
	}
	vals := []int{6, 2}

	for i := 0; i < len(keys); i++ {
		ans := trap(keys[i])
		if ans != vals[i] {
			fmt.Println("get", ans, "should be", vals[i], "case: #", i)
		}
	}
}
