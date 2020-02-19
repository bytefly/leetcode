package main

import (
	"fmt"
)

func selfDividingNumbers(left int, right int) []int {
	var ans []int
	for left <= right {
		n := left
		for n > 0 {
			t := n % 10
			if t == 0 || left%t != 0 {
				break
			}
			n /= 10
		}
		if n == 0 {
			ans = append(ans, left)
		}
		left++
	}
	return ans
}

func main() {
	fmt.Println(selfDividingNumbers(1, 22))
}
