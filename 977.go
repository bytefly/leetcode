package main

import (
	"fmt"
)

func sortedSquares(A []int) []int {
	ans := make([]int, len(A))
	left, right := 0, len(A)-1
	index := right
	for index >= 0 {
		lSquare, rSquare := A[left]*A[left], A[right]*A[right]
		if lSquare >= rSquare {
			ans[index] = lSquare
			left++
		} else {
			ans[index] = rSquare
			right--
		}
		index--
	}

	return ans
}

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 3, 10}))
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11}))
}
