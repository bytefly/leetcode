package main

import (
	"fmt"
)

func canPermutePalindrome(s string) bool {
	m := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	oddNum := 0
	for _, v := range m {
		if v%2 != 0 {
			oddNum++
			if oddNum > 1 {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Println(canPermutePalindrome("code"))
	fmt.Println(canPermutePalindrome("aab"))
	fmt.Println(canPermutePalindrome("carerac"))
}
