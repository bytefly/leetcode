package main

import (
	"fmt"
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	ss := s + s
	if strings.Index(ss[1:len(ss)-1], s) >= 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(repeatedSubstringPattern("ababaabb"))
	fmt.Println(repeatedSubstringPattern("abaaba"))
	fmt.Println(repeatedSubstringPattern("aabaab"))
	fmt.Println(repeatedSubstringPattern("abab"))
	fmt.Println(repeatedSubstringPattern("aba"))
	fmt.Println(repeatedSubstringPattern("abcabcabcabc"))
}
