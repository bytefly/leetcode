package main

import "fmt"

func sumOfDigits(A []int) int {
	min := A[0]
	for _, num := range A {
		if min > num {
			min = num
		}
	}
	var cnt int
	for min > 0 {
		cnt += min % 10
		min /= 10
	}

	return 1 - cnt%2
}

func main() {
	fmt.Println(sumOfDigits([]int{34, 23, 1, 24, 75, 33, 54, 8}))
	fmt.Println(sumOfDigits([]int{99, 77, 33, 66, 55}))
}
