package main

import (
	"fmt"
)

func isUgly(num int) bool {
	if num < 1 {
		return false
	}
	for {
		if num == 1 {
			break
		}
		t := num
		if num%2 == 0 {
			num >>= 1
		}
		if num%3 == 0 {
			num /= 3
		}
		if num%5 == 0 {
			num /= 5
		}
		if num == t {
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
