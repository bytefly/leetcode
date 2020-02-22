package main

import (
	"fmt"
)

func rotatedDigits(N int) int {
	var ans int
	m := [10]int{0, 1, 5, -1, -1, 2, 9, -1, 8, 6}
	for i := 1; i <= N; i++ {
		t, n := 0, i
		mul := 1
		for n > 0 {
			if m[n%10] < 0 {
				t = 0
				break
			}

			t += m[n%10] * mul
			n /= 10
			mul *= 10
		}

		if t != 0 && t != i {
			ans++
		}
	}

	return ans
}

func main() {
	fmt.Println(rotatedDigits(10))
	fmt.Println(rotatedDigits(1250))
}
