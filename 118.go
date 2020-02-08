package main

import (
	"fmt"
)

func generate(numRows int) [][]int {
	ans := make([][]int, numRows)
	if numRows < 1 {
		return ans
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

	return ans
}

func main() {
	ans := generate(8)
	for i := 0; i < len(ans); i++ {
		for j := 0; j < len(ans)-i; j++ {
			fmt.Print(" ")
		}
		fmt.Println(ans[i])
	}
}
