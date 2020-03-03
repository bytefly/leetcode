package main

import (
	"fmt"
	"strconv"
)

func isArmstrong(N int) bool {
	var sum int
	num, k := N, len(strconv.Itoa(N))
	for num > 0 {
		m, t := 1, num%10
		for i := 0; i < k; i++ {
			m *= t
		}
		sum += m
		num /= 10
	}

	return sum == N
}

func main() {
	fmt.Println(isArmstrong(153))
	fmt.Println(isArmstrong(123))
}
