package main

import "fmt"

func fixedPoint(A []int) int {
	left, right := 0, len(A)-1
	for left < right {
		mid := left + (right-left)/2
		if A[mid] >= mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if A[left] == left {
		return left
	}

	return -1
}

func main() {
	fmt.Println(fixedPoint([]int{-10}))
	fmt.Println(fixedPoint([]int{0, 2, 3, 4}))
	fmt.Println(fixedPoint([]int{-10, -5, 0, 3}))
	fmt.Println(fixedPoint([]int{-10, -5, 0, 2, 4}))
	fmt.Println(fixedPoint([]int{-10, -5, 0, 3, 7}))
	fmt.Println(fixedPoint([]int{0, 2, 5, 8, 17}))
	fmt.Println(fixedPoint([]int{-10, -5, 3, 4, 7, 9}))
	fmt.Println(fixedPoint([]int{-10, -5, 3, 4, 7, 9, 10}))
	fmt.Println(fixedPoint([]int{-10, -5, -2, 0, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(fixedPoint([]int{-10, -5, 2, 3, 7}))
}
