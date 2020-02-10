package main

import (
	"fmt"
)

func findMaxConsecutiveOnes(nums []int) int {
	var ans, count int

	for _, i := range nums {
		if i == 1 {
			count++
		} else {
			if count > ans {
				ans = count
			}
			count = 0
		}
	}

	if count > ans {
		ans = count
	}
	return ans
}

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{0}))
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 1, 1, 1, 1}))
	fmt.Println(findMaxConsecutiveOnes([]int{0, 0, 1, 1, 1, 1}))
	fmt.Println(findMaxConsecutiveOnes([]int{0, 0, 1, 1, 1, 0}))
	fmt.Println(findMaxConsecutiveOnes([]int{0, 0, 1, 1, 0, 1, 0}))
}
