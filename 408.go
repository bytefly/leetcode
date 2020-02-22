package main

import (
	"fmt"
	"strconv"
)

func validWordAbbreviation(word string, abbr string) bool {
	var wordPos int
	abbrPos := -1
	for i, v := range abbr {
		if v >= '0' && v <= '9' {
			if abbrPos < 0 {
				//invalid number started with 0s
				if v == '0' {
					return false
				}
				abbrPos = i
			}
		} else {
			if abbrPos >= 0 {
				n, _ := strconv.Atoi(abbr[abbrPos:i])
				//move word pos in the area
				wordPos += n
				abbrPos = -1
			}
			//compare character
			if wordPos == len(word) || word[wordPos] != byte(v) {
				return false
			}
			wordPos++
		}
	}
	//move word pos in the tail with abbr pos
	if abbrPos >= 0 {
		n, _ := strconv.Atoi(abbr[abbrPos:])
		if wordPos+n > len(word) {
			return false
		}
		wordPos += n
	}
	//word pos must in the tail
	if wordPos < len(word) {
		return false
	}

	return true
}

func main() {
	fmt.Println(validWordAbbreviation("internationalization", "i12iz4n"))
	fmt.Println(validWordAbbreviation("apple", "a2e"))
	fmt.Println(validWordAbbreviation("apple", "a4"))
	fmt.Println(validWordAbbreviation("apple", "4e"))
	fmt.Println(validWordAbbreviation("apple", "5"))
	fmt.Println(validWordAbbreviation("apple", "6"))
	fmt.Println(validWordAbbreviation("a", "01"))
	fmt.Println(validWordAbbreviation("hi", "2i"))
	fmt.Println(validWordAbbreviation("hi", "1"))
	fmt.Println(validWordAbbreviation("internationalization", "i5a11o1"))
}
