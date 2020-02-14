package main

import (
	"fmt"
)

func isPrime(num int) bool {
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func isUgly(num int) bool {
	if num < 1 {
		return false
	}
	if num == 1 {
		return true
	}
	for i := 2; i <= num; i++ {
		if isPrime(i) && num%i == 0 && i != 2 && i != 3 && i != 5 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isUgly(6))
	fmt.Println(isUgly(7))
	fmt.Println(isUgly(8))
	fmt.Println(isUgly(14))
	fmt.Println(isUgly(937351770))
}
