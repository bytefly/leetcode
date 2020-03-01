package main

import (
	"fmt"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	var ans [][]int

	sort.Ints(arr)
	min := arr[1] - arr[0]
	ans = append(ans, []int{arr[0], arr[1]})

	for i := 2; i < len(arr); i++ {
		t := arr[i] - arr[i-1]
		if t <= min {
			if t < min {
				ans = ans[:0]
				min = t
			}
			ans = append(ans, []int{arr[i-1], arr[i]})
		}
	}

	return ans
}

func main() {
	fmt.Println(minimumAbsDifference([]int{2, 1}))
	fmt.Println(minimumAbsDifference([]int{4, 2, 1, 3}))
	fmt.Println(minimumAbsDifference([]int{1, 3, 6, 10, 15}))
	fmt.Println(minimumAbsDifference([]int{3, 8, -10, 23, 19, -4, -14, 27}))
}
