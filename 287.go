package main

import "fmt"

func findDuplicate(nums []int) int {
	left, right := 1, len(nums)-1

	for left < right {
		mid := left + (right-left)>>1
		cnt := 0
		for _, num := range nums {
			if num <= mid {
				cnt++
			}
		}

		if cnt > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
	fmt.Println(findDuplicate([]int{3, 1, 3, 4, 2}))
}
