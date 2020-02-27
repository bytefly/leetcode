package main

import (
	"fmt"
)

func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	k %= m * n

	for i := 0; i < k; i++ {
		t := grid[m-1][n-1]
		for j := m - 1; j >= 0; j-- {
			for l := n - 1; l > 0; l-- {
				grid[j][l] = grid[j][l-1]
			}
			if j > 0 {
				grid[j][0] = grid[j-1][n-1]
			}
		}
		grid[0][0] = t
	}

	return grid
}

func main() {
	fmt.Println(shiftGrid([][]int{
		[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9},
	}, 1))
	fmt.Println(shiftGrid([][]int{
		[]int{3, 8, 1, 9}, []int{19, 7, 2, 5}, []int{4, 6, 11, 10}, []int{12, 0, 21, 13},
	}, 4))
	fmt.Println(shiftGrid([][]int{
		[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9},
	}, 9))
	fmt.Println(shiftGrid([][]int{
		[]int{1, 2, 3},
	}, 2))
}
