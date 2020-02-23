package main

import (
	"fmt"
)

func subtractProductAndSum(n int) int {
	sum, mul, num := 0, 1, n
	for num > 0 {
		t := num % 10
		sum += t
		mul *= t
		num /= 10
	}

	return mul - sum
}

func main() {
	fmt.Println(subtractProductAndSum(234))
	fmt.Println(subtractProductAndSum(4421))
}
