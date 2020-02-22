package main

import (
	"fmt"
)

func rotateString(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	if A == B {
		return true
	}

	size := len(A)
	m := []byte(A)
	for i := 0; i < size-1; i++ {
		t := m[0]
		for j := 0; j < size-1; j++ {
			m[j] = m[j+1]
		}
		m[size-1] = t

		if string(m) == B {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(rotateString("abcde", "cdeab"))
	fmt.Println(rotateString("abcde", "abced"))
	fmt.Println(rotateString("abccde", "bccdea"))
}
