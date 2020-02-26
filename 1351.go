package main

import "fmt"

func countNegatives(grid [][]int) int {
	var ans int
	for i := 0; i < len(grid); i++ {
		left, right := 0, len(grid[0])-1
		for left < right {
			mid := (left + right) / 2
			if grid[i][mid] >= 0 {
				left = mid + 1
			} else {
				right--
			}
		}
		if grid[i][left] < 0 {
			ans += len(grid[0]) - left
		}
	}
	return ans
}

func main() {
	fmt.Println(countNegatives([][]int{
		[]int{4, 3, 2, -1}, []int{3, 2, 1, -1}, []int{1, 1, -1, -2}, []int{-1, -1, -2, -3},
	}))
	fmt.Println(countNegatives([][]int{
		[]int{3, 2}, []int{1, 0},
	}))
	fmt.Println(countNegatives([][]int{
		[]int{1, -1}, []int{-1, -1},
	}))
	fmt.Println(countNegatives([][]int{
		[]int{-1},
	}))
}
