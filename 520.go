package main

import (
	"fmt"
)

func detectCapitalUse(word string) bool {
	capitalNum := 0
	firstPos := -1
	for i := 0; i < len(word); i++ {
		if word[i] >= 'A' && word[i] <= 'Z' {
			capitalNum++
			if capitalNum == 1 {
				firstPos = i
			}
		}
	}

	if capitalNum == 0 || capitalNum == len(word) {
		return true
	}
	if capitalNum == 1 && firstPos == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(detectCapitalUse("A"))
	fmt.Println(detectCapitalUse("abc"))
	fmt.Println(detectCapitalUse("USA"))
	fmt.Println(detectCapitalUse("Google"))

	fmt.Println(detectCapitalUse("BAD"))
	fmt.Println(detectCapitalUse("bAD"))
	fmt.Println(detectCapitalUse("baD"))
	fmt.Println(detectCapitalUse("bAd"))
	fmt.Println(detectCapitalUse("BaD"))
}
