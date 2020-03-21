package main

import "fmt"

func findLHS(nums []int) int {
	var max int
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	for k, v := range m {
		num, ok := m[k+1]
		if ok {
			t := num + v
			if t > max {
				max = t
			}
		}
	}

	return max
}

func main() {
	fmt.Println(findLHS([]int{1, 3, 2, 2, 5, 2, 3, 7}))
}
