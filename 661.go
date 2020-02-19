package main

import (
	"fmt"
)

/*
func imageSmoother(M [][]int) [][]int {
	var ans [][]int
	for i := 0; i < len(M); i++ {
        ans = append(ans, make([]int, len(M[0]))
    }
	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M[0]); j++ {
			row[i-1][j-1] += M[i][j]
			ans[i-1][j] = M[i][j]
			ans[i-1][j+1] += M[i][j]

			row[i][j-1] += M[i][j]
			ans[i][j] = M[i][j]
			ans[i][j+1] += M[i][j]

			row[i+1][j-1] += M[i][j]
			row[i+1][j] += M[i][j]
			row[i+1][j+1] += M[i][j]
		}
	}
	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M[0]); j++ {
            if i ==0 || i == len(M)-1
        }
	return ans
}
*/
func imageSmoother(M [][]int) [][]int {
	var ans [][]int
	for i := 0; i < len(M); i++ {
		row := make([]int, len(M[0]))
		for j := 0; j < len(M[0]); j++ {
			cnt := 0

			if i > 0 {
				if j > 0 {
					row[j] += M[i-1][j-1]
					cnt++
				}
				row[j] += M[i-1][j]
				cnt++
				if j < len(M[0])-1 {
					row[j] += M[i-1][j+1]
					cnt++
				}
			}

			if j > 0 {
				row[j] += M[i][j-1]
				cnt++
			}
			row[j] += M[i][j]
			cnt++
			if j < len(M[0])-1 {
				row[j] += M[i][j+1]
				cnt++
			}

			if i < len(M)-1 {
				if j > 0 {
					row[j] += M[i+1][j-1]
					cnt++
				}
				row[j] += M[i+1][j]
				cnt++
				if j < len(M[0])-1 {
					row[j] += M[i+1][j+1]
					cnt++
				}
			}

			row[j] /= cnt
		}
		ans = append(ans, row)
	}

	return ans
}

func main() {
	fmt.Println(imageSmoother([][]int{[]int{1, 1, 1}, []int{1, 0, 1}, []int{1, 1, 1}}))
}
