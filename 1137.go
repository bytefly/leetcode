package main

import "fmt"

func tribonacci(n int) int {
	nums := [3]int{0, 1, 1}
	if n < 3 {
		return nums[n]
	}
	for i := 3; i <= n; i++ {
		nums[0] = nums[0] + nums[1] + nums[2]
		nums[0], nums[1], nums[2] = nums[1], nums[2], nums[0]
	}
	return nums[2]
}

func main() {
	fmt.Println(tribonacci(4))
	fmt.Println(tribonacci(25))
}
