package main

import "fmt"

func isMonotonic(A []int) bool {
	if len(A) == 1 {
		return true
	}

	var prev int
	for i := 1; i < len(A); i++ {
		switch {
		case A[i] > A[i-1]:
			if prev < 0 {
				return false
			}
			if prev == 0 {
				prev = 1
			}
		case A[i] < A[i-1]:
			if prev > 0 {
				return false
			}
			if prev == 0 {
				prev = -1
			}
		}
	}

	return true
}

func main() {
	fmt.Println(isMonotonic([]int{2, 2, 2, 1, 4, 5}))
	fmt.Println(isMonotonic([]int{1, 2, 2, 3}))
	fmt.Println(isMonotonic([]int{6, 5, 4, 4}))
	fmt.Println(isMonotonic([]int{1, 3, 2}))
	fmt.Println(isMonotonic([]int{1, 2, 4, 5}))
	fmt.Println(isMonotonic([]int{1, 1, 1}))
}
