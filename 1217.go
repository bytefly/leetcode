package main

import "fmt"

func minCostToMoveChips(chips []int) int {
	var cnt [2]int
	for _, num := range chips {
		if num%2 == 0 {
			cnt[0]++
		} else {
			cnt[1]++
		}
	}

	if cnt[0] >= cnt[1] {
		return cnt[1]
	}
	return cnt[0]
}

func main() {
	fmt.Println(minCostToMoveChips([]int{1, 2, 3}))
	fmt.Println(minCostToMoveChips([]int{2, 2, 2, 3, 3}))
	fmt.Println(minCostToMoveChips([]int{2, 2, 2, 3, 3, 3}))
}
