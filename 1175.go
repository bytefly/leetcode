package main

import "fmt"

func numPrimeArrangements(n int) int {
	var primeCnt int
	isPrime := func(num int) bool {
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				return false
			}
		}
		return true
	}

	for i := 2; i <= n; i++ {
		if isPrime(i) {
			primeCnt++
		}
	}

	ans := 1
	for i := 1; i <= primeCnt || i <= n-primeCnt; i++ {
		if i <= primeCnt {
			ans *= i
		}
		if i <= n-primeCnt {
			ans *= i
		}
		if ans >= 1000000007 {
			ans %= 1000000007
		}
	}
	return ans
}

func main() {
	fmt.Println(numPrimeArrangements(1))
	fmt.Println(numPrimeArrangements(2))
	fmt.Println(numPrimeArrangements(5))
	fmt.Println(numPrimeArrangements(100))
}
