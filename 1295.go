package main

import (
	"fmt"
)

func findNumbers(nums []int) int {
	var ans int
	for _, num := range nums {
		cnt := 0
		for num > 0 {
			cnt++
			num /= 10
		}
		if cnt%2 == 0 {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(findNumbers([]int{12, 345, 2, 6, 7896}))
	fmt.Println(findNumbers([]int{555, 901, 482, 1771}))
}
