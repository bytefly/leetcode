package main

import (
	"fmt"
)

func trailingZeroes(n int) int {
	var ans int

	for n > 0 {
		ans += n / 5
		n /= 5
	}
	return ans
}

func main() {
	fmt.Println(trailingZeroes(3))
	fmt.Println(trailingZeroes(5))
	fmt.Println(trailingZeroes(11))
	fmt.Println(trailingZeroes(35))
	fmt.Println(trailingZeroes(2147483647))
}
