package main

import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {
	i, j := 0, int(math.Sqrt(float64(c)))

	for i <= j {
		if i*i+j*j == c {
			return true
		}
		if i*i+j*j > c {
			j--
		} else {
			i++
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
