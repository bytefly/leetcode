package main

import (
	"fmt"
	"math"
)

func countPrimes(n int) int {
	var ans int
	for i := 2; i < n; i++ {
		isPrime := true
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if j != i && i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(countPrimes(5))
	fmt.Println(countPrimes(100))
}
