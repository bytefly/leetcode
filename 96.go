package main

import "fmt"

func numTrees(n int) int {
	if n < 2 {
		return 1
	}

	ans := make([]int, n+1)
	ans[0], ans[1], ans[2] = 1, 1, 2

	for i := 3; i <= n; i++ {
		for j := 1; j <= i; j++ {
			t := ans[j-1] * ans[i-j]
			ans[i] += t
		}
	}
	return ans[n]
}

func main() {
	fmt.Println(numTrees(3))
}
