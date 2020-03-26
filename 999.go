package main

import "fmt"

func numRookCaptures(board [][]byte) int {
	var ans, m, n int
loop:
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == 'R' {
				m, n = i, j
				break loop
			}
		}
	}

	//left
	for i := n - 1; i >= 0; i-- {
		if board[m][i] == 'B' {
			break
		}
		if board[m][i] == 'p' {
			ans++
			break
		}
	}
	//right
	for i := n + 1; i < 8; i++ {
		if board[m][i] == 'B' {
			break
		}
		if board[m][i] == 'p' {
			ans++
			break
		}
	}
	//up
	for i := m - 1; i >= 0; i-- {
		if board[i][n] == 'B' {
			break
		}
		if board[i][n] == 'p' {
			ans++
			break
		}
	}
	//down
	for i := m + 1; i < 8; i++ {
		if board[i][n] == 'B' {
			break
		}
		if board[i][n] == 'p' {
			ans++
			break
		}
	}

	return ans
}

func main() {
	fmt.Println(numRookCaptures([][]byte{
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', 'p', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', 'R', '.', '.', '.', 'p'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', 'p', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
	}))
	fmt.Println(numRookCaptures([][]byte{
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
		[]byte{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
		[]byte{'.', 'p', 'B', 'R', 'B', 'p', '.', '.'},
		[]byte{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
		[]byte{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
	}))
	fmt.Println(numRookCaptures([][]byte{
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', 'p', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', 'p', '.', '.', '.', '.'},
		[]byte{'p', 'p', '.', 'R', '.', 'p', 'B', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', 'B', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', 'p', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
	}))
}
