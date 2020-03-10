package main

import "fmt"

func longestWord(words []string) string {
	var ans string
	var i, maxLen int

	m := make(map[string]int, len(words))
	for i := 0; i < len(words); i++ {
		m[words[i]] = i
	}

	for _, word := range words {
		for i = 1; i <= len(word); i++ {
			if _, ok := m[word[:i]]; !ok {
				break
			}
		}

		if i > len(word) && len(word) >= maxLen {
			if len(word) > maxLen || word < ans {
				maxLen = len(word)
				ans = word
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(longestWord([]string{"w", "wo", "wor", "worl", "world"}))
	fmt.Println(longestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"}))
}
