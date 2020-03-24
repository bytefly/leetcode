package main

import "fmt"

func countBits(num int) []int {
	ans := make([]int, num+1)
	for i := 0; i <= num; i++ {
		ans[i] = ans[i>>1] + i&1
	}
	return ans
}

func main() {
	fmt.Println(countBits(2))
	fmt.Println(countBits(5))
	fmt.Println(countBits(20))
}
