package main

import (
	"fmt"
)

func checkPerfectNumber(num int) bool {
	if num < 1 {
		return false
	}

	sum := 0
	for i := 1; i*i <= num; i++ {
		if num%i == 0 {
			sum += i
			//add another part except itself
			if i*i != num {
				sum += num / i
			}
		}
	}

	return sum-num == num
}

func main() {
	fmt.Println(checkPerfectNumber(28))
	fmt.Println(checkPerfectNumber(496))
	fmt.Println(checkPerfectNumber(8128))
	fmt.Println(checkPerfectNumber(137438691328))
	fmt.Println(checkPerfectNumber(270))
	fmt.Println(checkPerfectNumber(0))
	fmt.Println(checkPerfectNumber(1))
	fmt.Println(checkPerfectNumber(2))
	fmt.Println(checkPerfectNumber(99999993))
}
