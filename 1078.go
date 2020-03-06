package main

import (
	"fmt"
	"strings"
)

func findOcurrences(text string, first string, second string) []string {
	var ans []string
	words := strings.Split(text, " ")
	for i := 1; i < len(words)-1; i++ {
		if words[i] == second && words[i-1] == first {
			ans = append(ans, words[i+1])
		}
	}

	return ans
}

func main() {
	fmt.Println(findOcurrences("alice is a good girl she is a good student", "a", "good"))
	fmt.Println(findOcurrences("we will we will rock you", "we", "will"))
}
