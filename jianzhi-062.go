package main

import "fmt"

func lastRemaining(n int, m int) int {
	var ans int

	for i := 2; i <= n; i++ {
		ans = (ans + m) % i
	}
	return ans
}

func main() {
	fmt.Println(lastRemaining(5, 1))
	fmt.Println(lastRemaining(5, 3))
	fmt.Println(lastRemaining(10, 17))
}
