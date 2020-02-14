package main

import (
	"fmt"
)

func rob(nums []int) int {
	var ans int
	var evenLast, oddLast, skinLast int
	last := -2

	for i := 0; i < len(nums); i += 4 {
		evenSum, oddSum, skinSum := 0, 0, 0
		if last < i-1 {
			evenSum = nums[i]
			evenLast = i
			if i+2 < len(nums) {
				evenSum += nums[i+2]
				evenLast = i + 2
			}
			skinSum = nums[i]
			skinLast = i
			if i+3 < len(nums) {
				skinSum += nums[i+3]
				skinLast = i + 3
			}
		}
		if last < i && i+1 < len(nums) {
			oddSum = nums[i+1]
			oddLast = i + 1
			if i+3 < len(nums) {
				oddSum += nums[i+3]
				oddLast = i + 3
			}
		}

		max := evenSum
		last = evenLast
		if max < oddSum {
			max = oddSum
			last = oddLast
		}
		if max < skinSum {
			max = skinSum
			last = skinLast
		}

		if last == i+3 && last+1 < len(nums) && max < nums[last+1] {
			max = evenSum
			last = evenSum
		}

		ans += max
	}

	return ans
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
	fmt.Println(rob([]int{2, 1, 1, 2}))
	fmt.Println(rob([]int{1}))
	//error, should be 103
	fmt.Println(rob([]int{1, 3, 1, 3, 100}))
}
