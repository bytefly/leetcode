package main

import (
	"fmt"
	"sort"
)

func repeatedNTimes(A []int) int {
	sort.Ints(A)
	mid := len(A) / 2

	if A[mid-1] == A[mid-2] {
		return A[mid-1]
	}
	return A[mid]
}

func main() {
	fmt.Println(repeatedNTimes([]int{1, 2, 3, 3}))
	fmt.Println(repeatedNTimes([]int{2, 1, 2, 5, 3, 2}))
	fmt.Println(repeatedNTimes([]int{5, 1, 5, 2, 5, 3, 5, 4}))
	fmt.Println(repeatedNTimes([]int{9, 5, 3, 3}))
}
