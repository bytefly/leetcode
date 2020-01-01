package main

import (
	"fmt"
)

func maxArea(height []int) int {
	maxStart, maxEnd := -1, -1
	maxV := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			curHeight := height[i]
			if height[i] > height[j] {
				curHeight = height[j]
			}

			curV := (j - i) * curHeight
			if curV > maxV {
				maxV = curV
				maxStart, maxEnd = i, j
			}
		}
	}
	_, _ = maxStart, maxEnd
	fmt.Println(maxStart, maxEnd, maxV)
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
