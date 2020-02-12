package main

import (
	"fmt"
)

func isHappy(n int) bool {
	var m, t int
	if n < 1 {
		return false
	}
	for n != 1 {
		for n > 0 {
			m, n = n%10, n/10
			t += m * m
		}

		n, t = t, 0
		//compare with the smallest computed number
		if n == 4 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isHappy(1))
	fmt.Println(isHappy(2))
	fmt.Println(isHappy(3))
	fmt.Println(isHappy(4))
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(100))
}
