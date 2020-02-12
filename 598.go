package main

import (
	"fmt"
)

func maxCount(m int, n int, ops [][]int) int {
	minH, minV := m, n

	for _, t := range ops {
		if t[0] < minH {
			minH = t[0]
		}
		if t[1] < minV {
			minV = t[1]
		}
	}
	return minH * minV
}

func main() {
	fmt.Println(maxCount(3, 3, [][]int{}))
	fmt.Println(maxCount(3, 3, [][]int{[]int{1, 1}, []int{1, 2}}))
	fmt.Println(maxCount(3, 3, [][]int{[]int{2, 4}, []int{3, 5}}))
	fmt.Println(maxCount(3, 3, [][]int{[]int{2, 2}, []int{3, 3}}))
	fmt.Println(maxCount(10, 10, [][]int{[]int{2, 3}, []int{5, 7}}))
}
