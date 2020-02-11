package main

import (
	"fmt"
)

func islandPerimeter(grid [][]int) int {
	var ans int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				//up
				if i > 0 && grid[i-1][j] == 1 {
					ans++
				}
				//down
				if i < len(grid)-1 && grid[i+1][j] == 1 {
					ans++
				}
				//left
				if j > 0 && grid[i][j-1] == 1 {
					ans++
				}
				//right
				if j < len(grid[i])-1 && grid[i][j+1] == 1 {
					ans++
				}
			} else {
				//top
				if i == 0 {
					ans++
				}
                //bottom
				if i == len(grid)-1 {
					ans++
				}
				//leftmost
				if j == 0 {
					ans++
				}
                //rightmost
				if j == len(grid[i])-1 {
					ans++
				}
			}
		}
	}
	return ans
}

func main() {
	g := [][]int{[]int{0, 1, 0, 0},
		[]int{1, 1, 1, 0},
		[]int{0, 1, 0, 0},
		[]int{1, 1, 0, 0}}
	fmt.Println(islandPerimeter(g))
	g = [][]int{[]int{1}}
	fmt.Println(islandPerimeter(g))
}
