package main

import "fmt"

func binaryGap(N int) int {
	var i, ans int

	pos := -1
	for N > 0 {
		if N%2 == 1 {
			if pos >= 0 && i-pos > ans {
				ans = i - pos
			}
			pos = i
		}

		i++
		N /= 2
	}
	return ans
}

func main() {
	fmt.Println(binaryGap(22))
	fmt.Println(binaryGap(5))
	fmt.Println(binaryGap(6))
	fmt.Println(binaryGap(8))
}
