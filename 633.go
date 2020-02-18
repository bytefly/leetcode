package main

import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {
	for i := 0; i*i <= c; i++ {
		m := c - i*i
		j := int(math.Sqrt(float64(m)))
		if j*j == m {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(judgeSquareSum(0))
	fmt.Println(judgeSquareSum(5))
	fmt.Println(judgeSquareSum(3))
	fmt.Println(judgeSquareSum(13))
	fmt.Println(judgeSquareSum(25))
	fmt.Println(judgeSquareSum(2147483646))
}
