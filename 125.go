package main

import (
	"fmt"
)

func checkByte(b byte) byte {
	if b >= 'a' && b <= 'z' {
		return b
	}
	if b >= 'A' && b <= 'Z' {
		return b - 'A' + 'a'
	}
	if b >= '0' && b <= '9' {
		return b
	}
	return ' '
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	left, right := 0, len(s)-1
	for left < right {
		l, r := checkByte(s[left]), checkByte(s[right])
		if l == ' ' {
			left++
			continue
		}
		if r == ' ' {
			right--
			continue
		}
		if l != r {
			return false
		}
		left++
		right--
	}

	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome(""))
	fmt.Println(isPalindrome("1"))
	fmt.Println(isPalindrome("12"))
	fmt.Println(isPalindrome("121"))
	fmt.Println(isPalindrome("aa"))
}
