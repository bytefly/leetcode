package main

import (
	"fmt"
	"strconv"
)

//bubble sorting
func sort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

func findRelativeRanks(nums []int) []string {
	ans := make([]string, len(nums))
	scoreMap := make(map[int]int)
	for i, j := range nums {
		scoreMap[j] = i
	}

	sort(nums)

	for i := 0; i < len(ans); i++ {
		switch i {
		case 0:
			ans[scoreMap[nums[i]]] = "Gold Medal"
		case 1:
			ans[scoreMap[nums[i]]] = "Silver Medal"
		case 2:
			ans[scoreMap[nums[i]]] = "Bronze Medal"
		default:
			ans[scoreMap[nums[i]]] = strconv.FormatInt(int64(1+i), 10)
		}
	}

	return ans
}

func main() {
	fmt.Println(findRelativeRanks([]int{5, 4, 3, 2, 1}))
	fmt.Println(findRelativeRanks([]int{4, 2, 3, 1, 5}))
	fmt.Println(findRelativeRanks([]int{2, 4}))
}
