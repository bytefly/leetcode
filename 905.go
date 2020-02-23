package main

import (
	"fmt"
)

func sortArrayByParity(A []int) []int {
	ans := make([]int, len(A))
	left, right := 0, len(A)-1
	for i := 0; i < len(A); i++ {
		if A[i]%2 != 0 {
			ans[right] = A[i]
			right--
		} else {
			ans[left] = A[i]
			left++
		}
	}

	return ans
}

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
	fmt.Println(sortArrayByParity([]int{3, 1, 5, 7}))
	fmt.Println(sortArrayByParity([]int{2, 4, 6, 8}))
	fmt.Println(sortArrayByParity([]int{2}))
	fmt.Println(sortArrayByParity([]int{1}))
}
