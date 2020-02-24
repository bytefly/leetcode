package main

import (
	"fmt"
)

func tictactoe(moves [][]int) string {
	var m [3][3]byte

	for i := 0; i < len(moves); i++ {
		if i%2 == 0 {
			m[moves[i][0]][moves[i][1]] = 'A'
		} else {
			m[moves[i][0]][moves[i][1]] = 'B'
		}
	}

	for i := 0; i < 3; i++ {
		win := true
		for j := 1; j < 3; j++ {
			if m[i][j] != m[i][j-1] {
				win = false
				break
			}
		}
		if win && m[i][0] != 0 {
			return string(m[i][0])
		}
	}
	for i := 0; i < 3; i++ {
		win := true
		for j := 1; j < 3; j++ {
			if m[j][i] != m[j-1][i] {
				win = false
				break
			}
		}
		if win && m[0][i] != 0 {
			return string(m[0][i])
		}
	}

	if m[0][0] == m[1][1] && m[1][1] == m[2][2] && m[1][1] != 0 {
		if m[0][0] == 'A' {
			return "A"
		} else {
			return "B"
		}
	}
	if m[0][2] == m[1][1] && m[1][1] == m[2][0] && m[1][1] != 0 {
		if m[1][1] == 'A' {
			return "A"
		} else {
			return "B"
		}
	}

	if len(moves) == 9 {
		return "Draw"
	}
	return "Pending"
}

func main() {
	fmt.Println(tictactoe([][]int{[]int{0, 0}, []int{2, 0}, []int{1, 1}, []int{2, 1}, []int{2, 2}}))
	fmt.Println(tictactoe([][]int{[]int{0, 0}, []int{1, 1}, []int{0, 1}, []int{0, 2}, []int{1, 0}, []int{2, 0}}))
	fmt.Println(tictactoe([][]int{[]int{0, 0}, []int{1, 1}, []int{2, 0}, []int{1, 0}, []int{1, 2}, []int{2, 1}, []int{0, 1}, []int{0, 2}, []int{2, 2}}))
	fmt.Println(tictactoe([][]int{[]int{0, 0}, []int{1, 1}}))
	fmt.Println(tictactoe([][]int{[]int{0, 2}, []int{1, 0}, []int{2, 2}, []int{1, 2}, []int{2, 0}, []int{0, 0}, []int{0, 1}, []int{2, 1}, []int{1, 1}}))
}
