package main

import "fmt"

func prefixesDivBy5(A []int) []bool {
	var sum int
	ans := make([]bool, len(A))
	for i := 0; i < len(A); i++ {
		sum *= 2
		sum += A[i]
		sum %= 5
		ans[i] = sum == 0
	}

	return ans
}

func main() {
	fmt.Println(prefixesDivBy5([]int{0, 1, 1}))
	fmt.Println(prefixesDivBy5([]int{1, 1, 1}))
	fmt.Println(prefixesDivBy5([]int{0, 1, 1, 1, 1, 1}))
	fmt.Println(prefixesDivBy5([]int{1, 1, 1, 0, 1}))
	fmt.Println(prefixesDivBy5([]int{1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0}))
}
