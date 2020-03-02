package main

import (
	"fmt"
	"sort"
)

func heightChecker(heights []int) int {
	var ans int
	newHeights := make([]int, len(heights))
	copy(newHeights, heights)
	sort.Ints(newHeights)

	for i := 0; i < len(heights); i++ {
		if heights[i] != newHeights[i] {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(heightChecker([]int{1, 1, 4, 2, 1, 3}))
}
