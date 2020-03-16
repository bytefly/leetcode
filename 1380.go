package main

import "fmt"

func luckyNumbers(matrix [][]int) []int {
	var ans []int
	rowMap := make(map[int]bool, len(matrix))

	for i := 0; i < len(matrix); i++ {
		col, min := 0, matrix[i][0]
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] < min {
				col, min = j, matrix[i][j]
			}
		}
		rowMap[(i<<8)|col] = true
	}

	for i := 0; i < len(matrix[0]); i++ {
		row, max := 0, matrix[0][i]
		for j := 1; j < len(matrix); j++ {
			if matrix[j][i] > max {
				row, max = j, matrix[j][i]
			}
		}
		if rowMap[(row<<8)|i] == true {
			ans = append(ans, max)
		}
	}

	return ans
}

func main() {
	fmt.Println(luckyNumbers([][]int{
		[]int{3, 7, 8}, []int{9, 11, 13}, []int{15, 16, 17},
	}))
	fmt.Println(luckyNumbers([][]int{
		[]int{1, 10, 4, 2}, []int{9, 3, 8, 7}, []int{15, 16, 17, 12},
	}))
	fmt.Println(luckyNumbers([][]int{
		[]int{7, 8}, []int{1, 2},
	}))
}
