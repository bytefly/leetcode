package main

import (
	"fmt"
)

func backspaceCompare(S string, T string) bool {
	stackS, stackT := make([]byte, 0), make([]byte, 0)
	for i := 0; i < len(S); i++ {
		if S[i] == '#' {
			if len(stackS) > 0 {
				stackS = stackS[:len(stackS)-1]
			}
		} else {
			stackS = append(stackS, S[i])
		}
	}
	for i := 0; i < len(T); i++ {
		if T[i] == '#' {
			if len(stackT) > 0 {
				stackT = stackT[:len(stackT)-1]
			}
		} else {
			stackT = append(stackT, T[i])
		}
	}

	return string(stackS) == string(stackT)
}

func main() {
	fmt.Println(backspaceCompare("ab#c", "ad#c"))
	fmt.Println(backspaceCompare("ab##", "c#d#"))
	fmt.Println(backspaceCompare("a##c", "#a#c"))
	fmt.Println(backspaceCompare("a#c", "b"))
}
