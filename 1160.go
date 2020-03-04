package main

import "fmt"

func countCharacters(words []string, chars string) int {
	var charMap, charMapCpy [26]int
	var ans int

	for _, b := range chars {
		charMap[b-'a']++
	}
	charMapCpy = charMap

	for _, word := range words {
		rightWord := true
		for _, b := range word {
			if charMap[b-'a'] == 0 {
				rightWord = false
				break
			} else {
				charMap[b-'a']--
			}
		}
		if rightWord {
			ans += len(word)
		}
		charMap = charMapCpy
	}

	return ans
}

func main() {
	fmt.Println(countCharacters([]string{"cat", "bt", "hat", "tree"}, "atach"))
	fmt.Println(countCharacters([]string{"hello", "world", "leetcode"}, "welldonehoneyr"))
}
