package main

import (
	"fmt"
	"math"
)

func isPowerOfFour(num int) bool {
	if num < 1 {
		return false
	}
	if num&(num-1) != 0 {
		return false
	}

	t := int(math.Sqrt(float64(num)))
	if t*t != num || t&(t-1) != 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(isPowerOfFour(1))
	fmt.Println(isPowerOfFour(2))
	fmt.Println(isPowerOfFour(16))
	fmt.Println(isPowerOfFour(5))
}
