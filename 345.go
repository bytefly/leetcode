package main

import (
	"fmt"
)

func reverseVowels(s string) string {
	vowelTable := make([]int, 0)
	ans := []byte(s)
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' ||
			s[i] == 'A' || s[i] == 'E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U' {
			vowelTable = append(vowelTable, i)
		}
	}

	for i := 0; i < len(vowelTable)/2; i++ {
		t := ans[vowelTable[i]]
		ans[vowelTable[i]] = ans[vowelTable[len(vowelTable)-i-1]]
		ans[vowelTable[len(vowelTable)-i-1]] = t
	}
	return string(ans)
}

func main() {
	fmt.Println(reverseVowels("hello"))
	fmt.Println(reverseVowels("leetcode"))
	fmt.Println(reverseVowels("eo"))
}
