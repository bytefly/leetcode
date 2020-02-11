package main

import (
	"fmt"
)

func reverseWords(s string) string {
	var left, right int
	ans := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			ans[i] = s[i]
			if left < right {
				for m := left; m < right; m++ {
					ans[right-m+left-1] = s[m]
				}
				left = right
			}
		} else {
			if i == 0 || s[i-1] == ' ' {
				left, right = i, i+1
			} else {
				right++
			}
		}
	}

	if left < right {
		for m := left; m < right; m++ {
			ans[right-m+left-1] = s[m]
		}
	}

	return string(ans)
}

func main() {
	fmt.Println(reverseWords(""))
	fmt.Println(reverseWords("    "))
	fmt.Println(reverseWords("  e  "))
	fmt.Println(reverseWords("e"))
	fmt.Println(reverseWords("LeetCode"))
	fmt.Println(reverseWords("Let's take LeetCode contest"))
	fmt.Println(reverseWords("  Let's  take  LeetCode  contest  "))
}
