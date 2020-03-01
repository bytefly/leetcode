package main

import (
	"fmt"
	"sort"
)

func smallerNumbersThanCurrent(nums []int) []int {
	var ans []int
	newNums := make([]int, len(nums))
	copy(newNums, nums)
	sort.Ints(newNums)

	for _, num := range nums {
		left, right := 0, len(newNums)-1
		for left < right {
			mid := left + (right-left)/2
			if newNums[mid] >= num {
				right = mid
			} else {
				left = mid + 1
			}
		}

		ans = append(ans, left)
	}

	return ans
}

func main() {
	fmt.Println(smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3}))
	fmt.Println(smallerNumbersThanCurrent([]int{6, 5, 4, 8}))
	fmt.Println(smallerNumbersThanCurrent([]int{7, 7, 7, 7}))
}
