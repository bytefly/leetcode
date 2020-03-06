package main

import (
	"fmt"
	"sort"
)

func numEquivDominoPairs(dominoes [][]int) int {
	var ans int

	sort.SliceStable(dominoes, func(i, j int) bool {
		if dominoes[i][0] <= dominoes[j][0] {
			return true
		}
		return false
	})
	sort.SliceStable(dominoes, func(i, j int) bool {
		if dominoes[i][0] == dominoes[j][0] && dominoes[i][1] <= dominoes[j][1] {
			return true
		}
		return false
	})

	m := make(map[int]int, len(dominoes))
	for i, domino := range dominoes {
		m[domino[0]] = i
	}

	for i := 1; i <= len(dominoes); i++ {
		//multiple same [0] and [1]
		if i == len(dominoes) || dominoes[i-1][0] != dominoes[i][0] || dominoes[i-1][1] != dominoes[i][1] {
			n := i - 1
			for n > 0 && dominoes[n-1][0] == dominoes[n][0] && dominoes[n-1][1] == dominoes[n][1] {
				n--
			}
			ans += (i - n) * (i - n - 1) / 2
		}

		if i < len(dominoes) {
			n, ok := m[dominoes[i][1]]
			//multiple same cross [0][1]
			if ok && n < i {
				for n >= 0 && dominoes[n][0] == dominoes[i][1] {
					if dominoes[n][1] == dominoes[i][0] {
						ans++
					}
					n--
				}
			}
		}
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
