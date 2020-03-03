package main

import (
	"fmt"
	"sort"
)

func searchRight(nums []int, num int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] <= num {
			left = mid + 1
		} else {
			right = mid
		}
	}
	//not found
	if nums[left] < num {
		return 0
	}
	if nums[left] == num {
		return len(nums) - left - 1
	}
	return len(nums) - left
}

func numSmallerByFrequency(queries []string, words []string) []int {
	cntCharNum := func(s string) int {
		var cnt [26]int
		for _, b := range s {
			cnt[b-'a']++
		}
		for _, n := range cnt {
			if n > 0 {
				return n
			}
		}
		return 0
	}

	answer := make([]int, len(queries))
	wordCnt := make([]int, len(words))
	for i := 0; i < len(words); i++ {
		wordCnt[i] = cntCharNum(words[i])
	}
	sort.Ints(wordCnt)

	for i := 0; i < len(queries); i++ {
		answer[i] = searchRight(wordCnt, cntCharNum(queries[i]))
	}

	return answer
}

func main() {
	/*
		fmt.Println(searchRight([]int{0, 1, 1, 2, 3, 4, 4}, 1))
		fmt.Println(searchRight([]int{0, 1, 1, 2, 3, 4, 4}, 2))
		fmt.Println(searchRight([]int{0, 1, 1, 2, 3, 4, 4}, 4))
		fmt.Println(searchRight([]int{0, 1, 1, 2, 3, 4, 4}, 0))
		fmt.Println(searchRight([]int{0, 1, 1, 3, 4, 4}, 2))
		fmt.Println(searchRight([]int{0, 1, 1, 3, 4, 4}, 5))
	*/
	fmt.Println(numSmallerByFrequency([]string{"cbd"}, []string{"zaaaz"}))
	fmt.Println(numSmallerByFrequency([]string{"bbb", "cc"}, []string{"a", "aa", "aaa", "aaaa"}))
}
