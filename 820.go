package main

import (
	"fmt"
	"sort"
	"strings"
)

func minimumLengthEncoding(words []string) int {
	var m []string
	var ans int

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) >= len(words[j])
	})

	for _, word := range words {
		shortLen := 0
		for _, str := range m {
			if strings.HasSuffix(str, word) {
				shortLen = len(str)
			}
		}

		if shortLen == 0 {
			m = append(m, word)
			ans += len(word) + 1
		}
	}

	return ans
}

func main() {
	fmt.Println(minimumLengthEncoding([]string{"time", "me", "bell"}))
}
