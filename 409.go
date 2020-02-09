package main

import (
	"fmt"
)

func longestPalindrome(s string) int {
	var upperMap [26]int
	var lowerMap [26]int
	var ans int

	for _, j := range s {
		if j >= 'a' && j <= 'z' {
			lowerMap[j-'a']++
		} else {
			upperMap[j-'A']++
		}
	}

	for i := 0; i < 26; i++ {
		t := upperMap[i] / 2
		if t > 0 {
			ans += t * 2
		}
		t = lowerMap[i] / 2
		if t > 0 {
			ans += t * 2
		}
	}

	if ans < len(s) {
		ans++
	}
	return ans
}

func main() {
	fmt.Println(longestPalindrome("abccccdd"))
	fmt.Println(longestPalindrome("ccccdd"))
	fmt.Println(longestPalindrome("Aa"))
	fmt.Println(longestPalindrome(""))
}
