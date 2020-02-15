package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return true
		} else {
			m[num] = 1
		}
	}

	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1}))
	fmt.Println(containsDuplicate([]int{1, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
