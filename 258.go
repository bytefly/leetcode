package main

import (
	"fmt"
)

func addDigits(num int) int {
	for num > 9 {
		t := 0
		for num > 0 {
			t += num % 10
			num /= 10
		}
		num = t
	}
	return num
}

func main() {
	fmt.Println(addDigits(38))
}
