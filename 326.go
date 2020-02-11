package main

import (
	"fmt"
)

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}

	for n != 1 {
		if n%3 != 0 {
			return false
		}
		n /= 3
	}
	return true
}

func main() {
	fmt.Println(isPowerOfThree(0))
	fmt.Println(isPowerOfThree(1))
	fmt.Println(isPowerOfThree(3))
	fmt.Println(isPowerOfThree(27))
	fmt.Println(isPowerOfThree(45))
}
