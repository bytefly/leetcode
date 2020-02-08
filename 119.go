package main

import (
	"fmt"
)

func getRow(rowIndex int) []int {
	numRows := rowIndex + 1
	ans := make([][]int, numRows)
	if numRows < 1 {
		return []int{}
	}

	ans[0] = []int{1}
	for i := 1; i < numRows; i++ {
		row := make([]int, i+1)
		row[0], row[i] = 1, 1
		for j := 1; j < i; j++ {
			row[j] = ans[i-1][j-1] + ans[i-1][j]
		}
		ans[i] = row
	}

	return ans[rowIndex]
}

func main() {
	ans := getRow(0)
	fmt.Println(ans)
}
