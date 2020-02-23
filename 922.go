package main

import (
	"fmt"
)

func sortArrayByParityII(A []int) []int {
	left, right := 0, 1
	for left < len(A) && right < len(A) {
		if A[left]%2 == 0 {
			left += 2
			continue
		}
		if A[right]%2 != 0 {
			right += 2
			continue
		}
		A[left], A[right] = A[right], A[left]
		left += 2
		right += 2
	}

	return A
}

func main() {
	fmt.Println(sortArrayByParityII([]int{4, 3}))
	fmt.Println(sortArrayByParityII([]int{3, 4}))
	fmt.Println(sortArrayByParityII([]int{4, 2, 5, 7}))
	fmt.Println(sortArrayByParityII([]int{4, 5, 2, 7}))
	fmt.Println(sortArrayByParityII([]int{1, 2, 3, 4}))
}
