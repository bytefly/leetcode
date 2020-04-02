package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	var ans [][]int

	if len(intervals) == 0 {
		return nil
	}

	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0] && intervals[i][1] < intervals[j][1]
	})

	ans = append(ans, intervals[0])
	for i := 1; i < len(intervals); i++ {
		prev := ans[len(ans)-1]
		switch {
		case prev[1] >= intervals[i][0] && prev[1] < intervals[i][1]:
			prev[1] = intervals[i][1]
		case prev[1] < intervals[i][0]:
			ans = append(ans, intervals[i])
		}
	}

	return ans
}

func main() {
	fmt.Println(merge([][]int{
		[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18},
	}))
	fmt.Println(merge([][]int{
		[]int{1, 3}, []int{2, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18},
	}))
	fmt.Println(merge([][]int{
		[]int{1, 4}, []int{4, 5},
	}))
	fmt.Println(merge([][]int{
		[]int{1, 4},
	}))
	fmt.Println(merge([][]int{}))
}
