package main

import (
	"fmt"
)

func fib(N int) int {
	ans := []int{0, 1}

	for i := 2; i <= N; i++ {
		ans = append(ans, ans[i-1]+ans[i-2])
	}
	return ans[N]
}

func main() {
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
	fmt.Println(fib(30))
}
