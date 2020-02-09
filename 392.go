package main

import (
	"fmt"
)

func isSubsequence(s string, t string) bool {
	var m int
	for i, _ := range s {
		lastPos := -1
		for ; m < len(t); m++ {
			if t[m] == s[i] {
				lastPos = m
				break
			}
		}
		if lastPos < 0 {
			return false
		} else {
			m = lastPos + 1
		}
	}

	return true
}

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("abc", "ahgdcb"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))
	fmt.Println(isSubsequence("abbc", "ahgbcbdc"))
}
