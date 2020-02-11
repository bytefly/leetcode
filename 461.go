package main

import (
	"fmt"
)

func hammingDistance(x int, y int) int {
	var ans int
	z := x ^ y

	for i := 0; i <= 31; i++ {
		if z&(1<<uint(i)) != 0 {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(hammingDistance(1, 4))
}
