package main

import (
	"fmt"
)

func reverseStr(s string, k int) string {
	size := len(s)
	ans := make([]byte, size)
	left, right := 0, k
	for i := 0; i < size; i++ {
		if i == right {
			left = right + k
			right = left + k
		}
		if right > size {
			right = size
		}

		if i >= left && i < right {
			ans[right-i+left-1] = s[i]
		} else {
			ans[i] = s[i]
		}
	}
	return string(ans)
}

func main() {
	fmt.Println(reverseStr("abcdefg", 7))
	fmt.Println(reverseStr("abcdefg", 6))
	fmt.Println(reverseStr("abcdefg", 5))
	fmt.Println(reverseStr("abcdefg", 4))
	fmt.Println(reverseStr("abcdefg", 3))
	fmt.Println(reverseStr("abcdefg", 2))
	fmt.Println(reverseStr("abcdefg", 1))
	fmt.Println(reverseStr("abcdefghijklmnopqrstuvwxyz", 3))
}
