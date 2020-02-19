package main

import (
	"fmt"
)

func hasAlternatingBits(n int) bool {
	if n < 0 {
		return true
	}

	last := n & 1
	for n != 0 {
		n >>= 1
		cur := n & 1
		if last^cur != 1 {
			return false
		}
		last = cur
	}
	return true
}

func main() {
	fmt.Println(hasAlternatingBits(0))
	fmt.Println(hasAlternatingBits(1))
	fmt.Println(hasAlternatingBits(2))
	fmt.Println(hasAlternatingBits(3))
	fmt.Println(hasAlternatingBits(4))
	fmt.Println(hasAlternatingBits(5))
	fmt.Println(hasAlternatingBits(7))
	fmt.Println(hasAlternatingBits(10))
	fmt.Println(hasAlternatingBits(11))
}
