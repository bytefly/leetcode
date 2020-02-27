package main

import (
	"fmt"
)

func oddCells(n int, m int, indices [][]int) int {
	var oddRow, oddCol int
	rowMap := make(map[int]int, n)
	colMap := make(map[int]int, m)

	for _, indice := range indices {
		rowMap[indice[0]]++
		colMap[indice[1]]++
	}

	for _, v := range rowMap {
		if v%2 != 0 {
			oddRow++
		}
	}

	for _, v := range colMap {
		if v%2 != 0 {
			oddCol++
		}
	}

	return oddRow*(m-oddCol) + oddCol*(n-oddRow)
}

func main() {
	fmt.Println(oddCells(2, 3, [][]int{[]int{0, 1}, []int{1, 1}}))
	fmt.Println(oddCells(2, 2, [][]int{[]int{1, 1}, []int{0, 0}}))
}
