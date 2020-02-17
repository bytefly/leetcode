package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, str string) bool {
	tokens := strings.Split(str, " ")
	if len(tokens) != len(pattern) {
		return false
	}

	m := make(map[string]byte)
	n := make(map[byte]string)
	for i := 0; i < len(tokens); i++ {
		v, ok := m[tokens[i]]
		if ok {
			if v != pattern[i] {
				return false
			}
		} else {
			if _, ok = n[pattern[i]]; ok {
				return false
			}
			m[tokens[i]] = pattern[i]
			n[pattern[i]] = tokens[i]
		}
	}

	return true
}

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog cat cat fish"))
	fmt.Println(wordPattern("aaaa", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog dog dog dog"))
	fmt.Println(wordPattern("fbbf", "dog cat cat dog"))
}
