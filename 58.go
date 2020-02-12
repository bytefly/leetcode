package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
	var ans int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			ans++
		} else if ans > 0 {
			return ans
		}
	}

	return ans
}

func main() {
	fmt.Println(lengthOfLastWord("Hello world"))
	fmt.Println(lengthOfLastWord("a"))
	fmt.Println(lengthOfLastWord("a   "))
	fmt.Println(lengthOfLastWord("   a   "))
	fmt.Println(lengthOfLastWord("   a   b  "))
}
