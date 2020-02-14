package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var sCnt, tCnt [26]int
	for i := 0; i < len(s); i++ {
		sCnt[s[i]-'a']++
		tCnt[t[i]-'a']++
	}
	for i := 0; i < 26; i++ {
		if sCnt[i] != tCnt[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
}
