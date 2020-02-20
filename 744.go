package main

import (
	"fmt"
)

func nextGreatestLetter(letters []byte, target byte) byte {
	var ans, minGap byte
	minGap = 26

	for _, b := range letters {
		t := (b + 26 - target) % 26
		if t != 0 && t < minGap {
			minGap = t
			ans = b
		}
	}
	return ans
}

func main() {
	fmt.Printf("%c\n", nextGreatestLetter([]byte("cfj"), 'a'))
	fmt.Printf("%c\n", nextGreatestLetter([]byte("cfj"), 'c'))
	fmt.Printf("%c\n", nextGreatestLetter([]byte("cfj"), 'd'))
	fmt.Printf("%c\n", nextGreatestLetter([]byte("cfj"), 'g'))
	fmt.Printf("%c\n", nextGreatestLetter([]byte("ab"), 'z'))
	fmt.Printf("%c\n", nextGreatestLetter([]byte("cfj"), 'j'))
	fmt.Printf("%c\n", nextGreatestLetter([]byte("cfj"), 'k'))
}
