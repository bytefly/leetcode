package main

import "fmt"

func largestUniqueNumber(A []int) int {
	var cnts [1001]int
	for i := 0; i < len(A); i++ {
		cnts[A[i]]++
	}

	for i := 1000; i >= 0; i-- {
		if cnts[i] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(largestUniqueNumber([]int{5, 7, 3, 9, 4, 9, 8, 3, 1}))
	fmt.Println(largestUniqueNumber([]int{9, 9, 8, 8}))
}
