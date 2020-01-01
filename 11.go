package main

import (
	"fmt"
)

func maxArea(height []int) int {
	maxStart, maxEnd := -1, -1
	maxV := 0
	l, r := 0, len(height)-1
	for l < r {
		curHeight := height[l]
		curWidth := r - l
		if height[l] > height[r] {
			curHeight = height[r]
			r--
		} else {
			l++
		}

		curV := curWidth * curHeight
		if curV > maxV {
			maxV = curV
			maxStart, maxEnd = l, r
		}
	}
	_, _ = maxStart, maxEnd
	//fmt.Println(maxStart, maxEnd, maxV)
	return maxV
}

func main() {
	m := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	v := 49
	if v != m {
		fmt.Println("test fail, should be:", v, "but get", m)
	}

	m = maxArea([]int{1, 8})
	v = 1
	if v != m {
		fmt.Println("test fail, should be:", v, "but get", m)
	}
}
