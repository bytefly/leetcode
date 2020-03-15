package main

import (
	"fmt"
	"sort"
)

func makeNext(pattern string) []int {
	var j int
	ans := make([]int, len(pattern))

	//first char not matched but cannot walk backward
	ans[0] = 0
	for i := 1; i < len(pattern); {
		//prefix substring match
		if pattern[i] == pattern[j] {
			j++
			ans[i] = j
			i++
			continue
		}
		if j > 0 {
			//walk backward until matched or to the head
			j = ans[j-1]
		} else {
			//walk backward to the head
			ans[i] = 0
			i++
		}
	}

	return ans
}

func indexPairs(text string, words []string) [][]int {
	var ans [][]int

	for _, word := range words {
		next := makeNext(word)
		i, j := 0, 0
		for i < len(text) {
			if text[i] == word[j] {
				//matched
				if j == len(word)-1 {
					ans = append(ans, []int{i - j, i})
					j = next[j]
				} else {
					j++
				}
				i++
			} else {
				if j == 0 {
					i++
				} else {
					j = next[j-1]
				}
			}
		}
	}

	sort.SliceStable(ans, func(i, j int) bool {
		return ans[i][0] <= ans[j][0]
	})
	sort.SliceStable(ans, func(i, j int) bool {
		return ans[i][0] <= ans[j][0] && ans[i][1] < ans[j][1]
	})

	return ans
}

func main() {
	fmt.Println(indexPairs("thestoryofleetcodeandme", []string{"story", "fleet", "leetcode"}))
	fmt.Println(indexPairs("ababa", []string{"aba", "ab"}))
	fmt.Println(indexPairs("aaaabaabab", []string{"baa", "ba", "abab", "abb", "aa", "bb", "abba", "bbba", "ab", "aaba"}))
}
