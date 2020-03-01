package main

import (
	"fmt"
)

func countLetters(S string) int {
	var ans, prev int
	for i := 1; i < len(S); i++ {
		if S[i] != S[i-1] {
			n := i - prev
			ans += n * (n + 1) / 2
			prev = i
		}
	}
	n := len(S) - prev
	ans += n * (n + 1) / 2

	return ans
}

func main() {
	fmt.Println(countLetters("aa"))
	fmt.Println(countLetters("aaaba"))
	fmt.Println(countLetters("aaaaaaaaaa"))
}
