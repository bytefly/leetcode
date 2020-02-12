package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	valMap := make(map[int]int)
	var ans []int
	for i := 0; i < len(numbers); i++ {
		t := target - numbers[i]
		m, ok := valMap[t]
		if ok {
			if m > i {
				return []int{i + 1, m + 1}
			} else {
				return []int{m + 1, i + 1}
			}
		} else {
			valMap[numbers[i]] = i
		}
	}
	return ans
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 18))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 17))
	fmt.Println(twoSum([]int{2, 7, 11, 11, 11}, 22))
}
