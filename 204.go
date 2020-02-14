package main

import (
	"fmt"
)

func countPrimes(n int) int {
	var ans int
	primeTable := make([]bool, n)
	for i := 2; i < n; i++ {
		if !primeTable[i] {
			ans++
			for j := i + i; j < n; j += i {
				primeTable[j] = true
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(countPrimes(5))
	fmt.Println(countPrimes(100))
}
