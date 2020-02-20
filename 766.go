package main

import (
	"fmt"
)

func isToeplitzMatrix(matrix [][]int) bool {
	rowNum, colNum := len(matrix), len(matrix[0])

	for i := 1; i < rowNum; i++ {
		for j := 0; j < colNum; j++ {
			if j > 0 && matrix[i][j] != matrix[i-1][j-1] {
				return false
			}
		}
	}
	return true
}

func main() {
	m := [][]int{[]int{1, 2, 3, 4},
		[]int{5, 1, 2, 3},
		[]int{9, 5, 1, 2}}
	fmt.Println(isToeplitzMatrix(m))
}
