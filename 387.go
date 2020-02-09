package main

import (
	"fmt"
)

func firstUniqChar(s string) int {
	var charNum [26]int
	var uniqCharPos [26]int
	firstPos := len(s) + 1

	for i, j := range s {
		charNum[j-'a']++
		if charNum[j-'a'] > 1 {
			// mark ununiq char
			uniqCharPos[j-'a'] = len(s) + 1
		} else {
			// mark uniq char position
			uniqCharPos[j-'a'] = i + 1
		}
	}
	// find the minimum position in all uniq chars
	for _, j := range uniqCharPos {
		if j < firstPos && j > 0 {
			firstPos = j
		}
	}
	// not found
	if firstPos > len(s) {
		return -1
	}
	return firstPos - 1
}

func main() {
	fmt.Println(firstUniqChar(""))
	fmt.Println(firstUniqChar("a"))
	fmt.Println(firstUniqChar("aa"))
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))
}
