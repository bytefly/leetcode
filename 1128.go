package main

import (
	"fmt"
)

func numEquivDominoPairs(dominoes [][]int) int {
	var ans, key int

	m := make(map[int]int, len(dominoes))
	for _, domino := range dominoes {
		if domino[0] <= domino[1] {
			key = domino[0]*10 + domino[1]
		} else {
			key = domino[1]*10 + domino[0]
		}
		m[key]++
	}
	for _, v := range m {
		ans += v * (v - 1) / 2
	}

	return ans
}

func main() {
	fmt.Println(numEquivDominoPairs([][]int{
		[]int{1, 2}, []int{2, 1}, []int{3, 4}, []int{5, 6},
	}))
	fmt.Println(numEquivDominoPairs([][]int{
		[]int{1, 2}, []int{2, 1}, []int{2, 8}, []int{3, 4}, []int{8, 2}, []int{5, 6},
	}))
	fmt.Println(numEquivDominoPairs([][]int{
		[]int{1, 2}, []int{1, 2}, []int{1, 1}, []int{1, 2}, []int{2, 2},
	}))
	fmt.Println(numEquivDominoPairs([][]int{
		[]int{1, 2}, []int{2, 1}, []int{2, 8}, []int{3, 4}, []int{8, 2}, []int{5, 6}, []int{2, 8}, []int{8, 2},
	}))
	fmt.Println(numEquivDominoPairs([][]int{
		[]int{1, 1}, []int{2, 2}, []int{1, 1}, []int{1, 2}, []int{1, 2}, []int{1, 1},
	}))
	fmt.Println(numEquivDominoPairs([][]int{
		[]int{2, 2}, []int{1, 2}, []int{1, 2}, []int{1, 1}, []int{1, 2}, []int{1, 1}, []int{2, 2},
	}))
	fmt.Println(numEquivDominoPairs([][]int{
		[]int{2, 2}, []int{1, 2}, []int{2, 1}, []int{2, 1}, []int{1, 1}, []int{1, 2}, []int{2, 2}, []int{1, 1}, []int{1, 1}, []int{1, 1}, []int{2, 2}, []int{1, 2}, []int{1, 1}, []int{2, 1}, []int{2, 1}, []int{2, 1}, []int{2, 2}, []int{1, 2}, []int{2, 2}, []int{2, 1},
	}))
	fmt.Println(numEquivDominoPairs([][]int{
		[]int{2, 1}, []int{5, 4}, []int{3, 7}, []int{6, 2}, []int{4, 4}, []int{1, 8}, []int{9, 6}, []int{5, 3}, []int{7, 4}, []int{1, 9}, []int{1, 1}, []int{6, 6}, []int{9, 6}, []int{1, 3}, []int{9, 7}, []int{4, 7}, []int{5, 1}, []int{6, 5}, []int{1, 6}, []int{6, 1}, []int{1, 8}, []int{7, 2}, []int{2, 4}, []int{1, 6}, []int{3, 1}, []int{3, 9}, []int{3, 7}, []int{9, 1}, []int{1, 9}, []int{8, 9},
	}))
}
