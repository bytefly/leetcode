package main

import (
	"fmt"
	"sort"
)

func maxNumberOfApples(arr []int) int {
	var sum, ans int
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if sum <= 5000 {
			ans++
		} else {
			break
		}
	}

	return ans
}

func main() {
	fmt.Println(maxNumberOfApples([]int{100, 200, 150, 1000}))
	fmt.Println(maxNumberOfApples([]int{900, 950, 800, 1000, 700, 800}))
}
