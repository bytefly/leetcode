package main

import (
	"fmt"
)

func getSum(a int, b int) int {
	m, n := a^b, a&b
	//loop until carry is 0
	for n != 0 {
		a, b = m, n<<1
		m, n = a^b, a&b
	}

	return m
}

func main() {
	fmt.Println(getSum(1, 2))
	fmt.Println(getSum(-2, 3))
	fmt.Println(getSum(589, 399))
}
