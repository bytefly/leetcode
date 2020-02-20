package main

import (
	"fmt"
)

func shortestCompletingWord(licensePlate string, words []string) string {
	var plateFP, wordFP [26]byte
	var matchWord string
	for _, c := range licensePlate {
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			plateFP[c|32-'a']++
		}
	}
	for _, word := range words {
		for _, c := range word {
			if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
				wordFP[c|32-'a']++
			}
		}

		found := true
		for i := 0; i < 26; i++ {
			if wordFP[i] < plateFP[i] {
				found = false
			}
			wordFP[i] = 0
		}
		if found && (matchWord == "" || len(matchWord) > len(word)) {
			matchWord = word
		}
	}

	return matchWord
}

func main() {
	fmt.Println(shortestCompletingWord("1s3 PSt", []string{"step", "steps", "striple", "stepple"}))
	fmt.Println(shortestCompletingWord("1s3 456", []string{"looks", "pest", "stew", "show"}))
}
