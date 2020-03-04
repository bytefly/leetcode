package main

import (
	"fmt"
	"sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	var ans, head []int
	m := make(map[int]int, len(arr2))
	for _, num := range arr2 {
		m[num] = 0
	}
	sort.Ints(arr1)

	for _, num := range arr1 {
		if _, ok := m[num]; ok {
			m[num]++
		} else {
			ans = append(ans, num)
		}
	}

	for _, num := range arr2 {
		for i := 0; i < m[num]; i++ {
			head = append(head, num)
		}
	}

	ans = append(head, ans...)
	return ans
}

func main() {
	fmt.Println(relativeSortArray([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}))
}
