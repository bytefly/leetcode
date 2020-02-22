package main

import (
	"fmt"
)

func isDigit(c byte) bool {
	if c >= 'a' && c <= 'z' {
		return true
	}
	if c >= 'A' && c <= 'Z' {
		return true
	}
	return false
}

func reverseOnlyLetters(S string) string {
	ans := []byte(S)
	left, right := 0, len(ans)-1
	for {
		for left < len(ans) && !isDigit(ans[left]) {
			left++
		}
		for right >= 0 && !isDigit(ans[right]) {
			right--
		}

		if left >= right {
			break
		}

		ans[left], ans[right] = ans[right], ans[left]
		left++
		right--
	}

	return string(ans)
}

func main() {
	fmt.Println(reverseOnlyLetters("?6C40E"))
	fmt.Println(reverseOnlyLetters("7_A28]"))
	fmt.Println(reverseOnlyLetters("7_28]"))
	fmt.Println(reverseOnlyLetters("ab-d"))
	fmt.Println(reverseOnlyLetters("ab-cd"))
	fmt.Println(reverseOnlyLetters("a-bC-dEf-ghIj"))
	fmt.Println(reverseOnlyLetters("Test1ng-Leet=code-Q!"))
}
