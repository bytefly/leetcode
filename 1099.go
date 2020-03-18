package main

import (
	"fmt"
	"sort"
)

func twoSumLessThanK(A []int, K int) int {
	var left, right int

	ans := -1
	sort.Ints(A)
	if len(A) < 2 || A[0]+A[1] >= K {
		return ans
	}
	for right = len(A) - 1; right >= 0; right-- {
		if A[right] < K {
			break
		}
	}

	for left < right {
		t := A[left] + A[right]
		if t >= K {
			right--
		} else {
			if t > ans {
				ans = t
			}
			left++
		}
	}

	return ans
}

func main() {
	fmt.Println(twoSumLessThanK([]int{34, 23, 1, 24, 75, 33, 54, 8}, 60))
	fmt.Println(twoSumLessThanK([]int{10, 20, 30}, 15))
}
