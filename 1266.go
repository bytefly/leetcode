package main

import (
	"fmt"
	"math"
)

func minTimeToVisitAllPoints(points [][]int) int {
	var ans int
	for i := 1; i < len(points); i++ {
		m, n := int(math.Abs(float64(points[i][0]-points[i-1][0]))), int(math.Abs(float64(points[i][1]-points[i-1][1])))
		if m >= n {
			ans += m
		} else {
			ans += n
		}
	}
	return ans
}

func main() {
	fmt.Println(minTimeToVisitAllPoints([][]int{
		[]int{1, 1}, []int{3, 4}, []int{-1, 0},
	}))
	fmt.Println(minTimeToVisitAllPoints([][]int{
		[]int{3, 2}, []int{-2, 2},
	}))
	fmt.Println(minTimeToVisitAllPoints([][]int{
		[]int{3, 2},
	}))
}
