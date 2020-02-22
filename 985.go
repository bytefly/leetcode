package main

import (
	"fmt"
)

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	var sum int
	ans := make([]int, len(queries))

	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			sum += A[i]
		}
	}
	for i := 0; i < len(queries); i++ {
		if A[queries[i][1]]%2 == 0 {
			if queries[i][0]%2 == 0 {
				sum += queries[i][0]
			} else {
				sum -= A[queries[i][1]]
			}
		} else if queries[i][0]%2 != 0 {
			sum += queries[i][0] + A[queries[i][1]]
		}

		ans[i] = sum
		A[queries[i][1]] += queries[i][0]
	}

	return ans
}

func main() {
	fmt.Println(sumEvenAfterQueries([]int{1, 2, 3, 4}, [][]int{[]int{1, 0}, []int{-3, 1}, []int{-4, 0}, []int{2, 3}}))
}
