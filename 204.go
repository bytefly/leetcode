package main

import (
	"fmt"
)

func countPrimes(n int) int {
	var ans int
	primeTable := make([]byte, n)
	for i := 2; i < n; i++ {
		if 1 != primeTable[i] {
			ans++
			for j := i + i; j < n; j += i {
				primeTable[j] = 1
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(countPrimes(5))
	fmt.Println(countPrimes(100))
}
