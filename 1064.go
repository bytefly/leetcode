package main

import "fmt"

func fixedPoint(A []int) int {
	for i := 0; i < len(A); i++ {
		if A[i] == i {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(fixedPoint([]int{-10, -5, 0, 3, 7}))
	fmt.Println(fixedPoint([]int{0, 2, 5, 8, 17}))
	fmt.Println(fixedPoint([]int{-10, -5, 3, 4, 7, 9}))
}
