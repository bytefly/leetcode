package main

import (
	"fmt"
)

func minDeletionSize(A []string) int {
	var ans int
	for i := 0; i < len(A[0]); i++ {
		for j := 1; j < len(A); j++ {
			if A[j][i] < A[j-1][i] {
				ans++
				break
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(minDeletionSize([]string{"cba", "daf", "ghi"}))
}
