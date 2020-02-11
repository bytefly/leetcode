package main

import (
	"fmt"
)

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	if n&(n-1) == 0 {
		return true
	}

	return false
}

func main() {
	fmt.Println(isPowerOfTwo(0))
	fmt.Println(isPowerOfTwo(1))
	fmt.Println(isPowerOfTwo(16))
	fmt.Println(isPowerOfTwo(218))
}
