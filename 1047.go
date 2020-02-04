package main

import (
	"fmt"
)

func removeDuplicates(S string) string {
	stack := make([]byte, 0)
	for i := 0; i < len(S); i++ {
		if len(stack) == 0 || stack[len(stack)-1] != S[i] {
			stack = append(stack, S[i])
			continue
		}
		stack = stack[:len(stack)-1]
	}
	return string(stack)
}

func main() {
	fmt.Println(removeDuplicates("abbaca"))
	fmt.Println(removeDuplicates("abbbaca"))
}
