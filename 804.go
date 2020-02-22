package main

import (
	"fmt"
)

func uniqueMorseRepresentations(words []string) int {
	table := [26]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	m := make(map[string]bool, len(words))
	for _, s := range words {
		mos := ""
		for _, b := range s {
			mos += table[b-'a']
		}
		m[mos] = true
	}

	return len(m)
}

func main() {
	fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
}
