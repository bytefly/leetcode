package main

import (
	"fmt"
	"sort"
)

func highFive(items [][]int) [][]int {
	var m [1000][]int
	var ans [][]int
	for i := 0; i < len(items); i++ {
		m[items[i][0]] = append(m[items[i][0]], items[i][1])
	}

	for i := 0; i < 1000; i++ {
		size := len(m[i])
		if size >= 5 {
			sort.Ints(m[i])
			ans = append(ans, []int{i, (m[i][size-1] + m[i][size-2] + m[i][size-3] + m[i][size-4] + m[i][size-5]) / 5})
		}
	}

	return ans
}

func main() {
	fmt.Println(highFive([][]int{
		[]int{1, 91}, []int{1, 92}, []int{2, 93}, []int{2, 97}, []int{1, 60}, []int{2, 77}, []int{1, 65}, []int{1, 87}, []int{1, 100}, []int{2, 100}, []int{2, 76},
	}))
}
