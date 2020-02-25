package main

import (
	"fmt"
)

func checkIfExist(arr []int) bool {
	var zeroCnt int
	m := make(map[int]int, len(arr))
	for i, num := range arr {
		if num == 0 {
			zeroCnt++
		} else {
			m[num] = i
		}
	}

	if zeroCnt > 1 {
		return true
	}
	for _, num := range arr {
		if _, ok := m[num*2]; ok {
			return true
		}
		if num%2 != 0 {
			continue
		}
		if _, ok := m[num*2]; ok {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(checkIfExist([]int{10, 2, 5, 3}))
	fmt.Println(checkIfExist([]int{7, 1, 14, 11}))
	fmt.Println(checkIfExist([]int{3, 1, 7, 11}))
	fmt.Println(checkIfExist([]int{3, 1, 7, 0}))
	fmt.Println(checkIfExist([]int{3, 0, 7, 0}))
}
