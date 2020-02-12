package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		t := digits[i] + carry
		if t > 9 {
			carry = 1
			digits[i] = t - 10
		} else {
			carry = 0
			digits[i] = t
		}
	}
	if carry > 0 {
		digits = append([]int{carry}, digits...)
	}

	return digits
}

func main() {
	fmt.Println(plusOne([]int{0}))
	fmt.Println(plusOne([]int{9, 9, 9}))
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
}
