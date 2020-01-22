package main

import (
	"fmt"
)

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}

	l, r := 1, len(height)-2
	lmax, rmax := height[l-1], height[r+1]
	ans := 0

	for l <= r {
		lmax = max(lmax, height[l])
		rmax = max(rmax, height[r])
		if lmax < rmax {
			ans += lmax - height[l]
			l++
		} else {
			ans += rmax - height[r]
			r--
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
