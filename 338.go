package main

import "fmt"

func countBits(num int) []int {
	ans := make([]int, num+1)
	for i := 0; i <= num; i++ {
		n, cnt := i, 0
		for n > 0 {
			if n&1 == 1 {
				cnt++
			}
			n >>= 1
		}
		ans[i] = cnt
	}
	return ans
}

func main() {
	fmt.Println(countBits(2))
	fmt.Println(countBits(5))
}
