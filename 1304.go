package main

import "fmt"

func sumZero(n int) []int {
	var j int
	ans := make([]int, n)
	if n%2 != 0 {
		j = 1
	}
	for i := 1; i <= n/2; i++ {
		ans[j], ans[j+1] = -i, i
		j += 2
	}
	return ans
}

func main() {
	fmt.Println(sumZero(4))
	fmt.Println(sumZero(2))
	fmt.Println(sumZero(5))
	fmt.Println(sumZero(3))
	fmt.Println(sumZero(1))
}
