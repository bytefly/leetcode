package main

import (
	"fmt"
	"sort"
)

func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int, len(arr))
	for _, num := range arr {
		m[num]++
	}

	cnt := make([]int, len(m))
	var i int
	for _, v := range m {
		cnt[i] = v
		i++
	}

	sort.Ints(cnt)
	for i := 1; i < len(cnt); i++ {
		if cnt[i] == cnt[i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
	fmt.Println(uniqueOccurrences([]int{1, 2}))
	fmt.Println(uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}))
}
