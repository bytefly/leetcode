package main

import (
	"fmt"
	"sort"
)

type cell struct {
	x, y int
	dist int
}

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	var ans [][]int
	var cells []cell

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			dist := 0
			if r0 >= i {
				dist += r0 - i
			} else {
				dist += i - r0
			}
			if c0 >= j {
				dist += c0 - j
			} else {
				dist += j - c0
			}
			cells = append(cells, cell{i, j, dist})
		}
	}

	sort.Slice(cells, func(i, j int) bool {
		return cells[i].dist <= cells[j].dist
	})

	for _, cord := range cells {
		ans = append(ans, []int{cord.x, cord.y})
	}
	return ans
}

func main() {
	fmt.Println(allCellsDistOrder(1, 2, 0, 0))
	fmt.Println(allCellsDistOrder(2, 2, 0, 1))
	fmt.Println(allCellsDistOrder(2, 3, 1, 2))
}
