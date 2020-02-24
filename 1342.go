package main

import (
	"fmt"
)

func numberOfSteps(num int) int {
	var ans int

	for num > 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}
		ans++
	}
	return ans
}

func main() {
	fmt.Println(numberOfSteps(0))
	fmt.Println(numberOfSteps(14))
	fmt.Println(numberOfSteps(8))
	fmt.Println(numberOfSteps(123))
}
