package main

import (
	"fmt"
)

func trailingZeroes(n int) int {
	var ans int

	for i := 0; i <= n; i += 5 {
		if i != 0 && i%5 == 0 {
			t := i
			for t >= 5 && t%5 == 0 {
				ans++
				t /= 5
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(trailingZeroes(3))
	fmt.Println(trailingZeroes(5))
	fmt.Println(trailingZeroes(11))
	fmt.Println(trailingZeroes(35))
}
