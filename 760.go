package main

import (
	"fmt"
)

func anagramMappings(A []int, B []int) []int {
	ans := make([]int, len(A))
	m := make(map[int]int, len(A))
	for i, j := range B {
		m[j] = i
	}
	for i, j := range A {
		ans[i] = m[j]
	}

	return ans
}

func main() {
	fmt.Println(anagramMappings([]int{12, 28, 46, 32, 50}, []int{50, 12, 32, 46, 28}))
}
