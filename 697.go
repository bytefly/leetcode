package main

import (
	"fmt"
)

type MetaData struct {
	Cnt      int
	FirstPos int
	LastPos  int
}

func findShortestSubArray(nums []int) int {
	hashMap := make(map[int]*MetaData, len(nums))
	ans, max := len(nums), 1

	for i := 0; i < len(nums); i++ {
		m, ok := hashMap[nums[i]]
		if ok {
			m.Cnt++
			m.LastPos = i
			if m.Cnt > max {
				max = m.Cnt
			}
		} else {
			hashMap[nums[i]] = &MetaData{1, i, i}
		}
	}

	for _, v := range hashMap {
		if v.Cnt == max {
			length := v.LastPos - v.FirstPos + 1
			if length < ans {
				ans = length
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(findShortestSubArray([]int{1}))
	fmt.Println(findShortestSubArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1}))
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2}))
}
