package main

import (
	"fmt"
)

func shortestDistance(words []string, word1 string, word2 string) int {
	prev, max := -1, len(words)
	for i, word := range words {
		if word == word1 || word == word2 {
			if prev >= 0 && words[prev] != word && i-prev < max {
				max = i - prev
			}
			prev = i
		}
	}

	return max
}

func main() {
	fmt.Println(shortestDistance([]string{"practice", "makes", "perfect", "coding", "makes"}, "coding", "practice"))
	fmt.Println(shortestDistance([]string{"practice", "makes", "perfect", "coding", "makes"}, "makes", "coding"))
	fmt.Println(shortestDistance([]string{"practice", "makes", "perfect", "coding", "makes"}, "makes", "practice"))
}
