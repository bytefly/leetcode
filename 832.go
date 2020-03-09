package main

import "fmt"

func flipAndInvertImage(A [][]int) [][]int {
	oddCols := len(A[0])%2 == 1
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[0])/2; j++ {
			A[i][j], A[i][len(A[0])-j-1] = A[i][len(A[0])-j-1], A[i][j]
			A[i][j] = 1 - A[i][j]
			A[i][len(A[0])-j-1] = 1 - A[i][len(A[0])-j-1]
		}
		if oddCols {
			A[i][len(A[0])/2] = 1 - A[i][len(A[0])/2]
		}
	}
	return A
}

func main() {
	fmt.Println(flipAndInvertImage([][]int{
		[]int{1, 1, 0}, []int{1, 0, 1}, []int{0, 0, 0},
	}))
	fmt.Println(flipAndInvertImage([][]int{
		[]int{1, 1, 0, 0}, []int{1, 0, 0, 1}, []int{0, 1, 1, 1}, []int{1, 0, 1, 0},
	}))
}
