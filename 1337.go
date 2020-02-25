package main

import (
	"fmt"
)

func kWeakestRows(mat [][]int, k int) []int {
	var ans []int
	m := make(map[int]bool, len(mat))
	for j := 0; j < len(mat[0]); j++ {
		for i := 0; i < len(mat); i++ {
			//add zero position
			if mat[i][j] == 0 {
				//only add new zero position
				if _, ok := m[i]; !ok {
					ans = append(ans, i)
					m[i] = true
				}
			}
		}
	}
	//fill ones' positions
	if len(ans) < len(mat) {
		for i := 0; i < len(mat); i++ {
			//key i not in map is one
			if _, ok := m[i]; !ok {
				ans = append(ans, i)
			}
		}
	}
	return ans[:k]
}

func main() {
	fmt.Println(kWeakestRows([][]int{
		[]int{1, 1, 0, 0, 0},
		[]int{1, 1, 1, 1, 0},
		[]int{1, 0, 0, 0, 0},
		[]int{1, 1, 0, 0, 0},
		[]int{1, 1, 1, 1, 1},
	}, 3))
	fmt.Println(kWeakestRows([][]int{
		[]int{1, 0, 0, 0},
		[]int{1, 1, 1, 1},
		[]int{1, 0, 0, 0},
		[]int{1, 0, 0, 0},
	}, 2))
}
