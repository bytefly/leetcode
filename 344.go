package main

import (
	"fmt"
)

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		t := s[i]
		s[i] = s[len(s)-i-1]
		s[len(s)-i-1] = t
	}
}

func main() {
	s := []byte("hello")
	reverseString(s)
	fmt.Println(string(s))

	s = []byte("byte")
	reverseString(s)
	fmt.Println(string(s))
}
