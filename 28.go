package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:len(needle)+i] == needle {
			return i
		}
	}
	return -1
}

func main() {
	strs := []string{"hello", "ll", "aaaaa", "bba", "", "", "a", "a"}
	res := []int{2, -1, 0, 0}

	for i := 0; i < len(res); i++ {
		m := strStr(strs[2*i], strs[2*i+1])
		if m != res[i] {
			fmt.Println("%d tests fail, shoulde be %d but %d", i, res[i], m)
		}
	}
}
