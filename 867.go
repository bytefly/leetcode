package main

import "fmt"

func transpose(A [][]int) [][]int {
	ans := make([][]int, len(A[0]))
	for i := 0; i < len(A[0]); i++ {
		ans[i] = make([]int, len(A))
		for j := 0; j < len(A); j++ {
			ans[i][j] = A[j][i]
		}
	}
	return ans
}

func main() {
	fmt.Println(transpose([][]int{
		[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9},
	}))
	fmt.Println(transpose([][]int{
		[]int{1, 2, 3}, []int{4, 5, 6},
	}))
}
