package main

import (
	"fmt"
)

func isValid(s string) bool {
	stack := make([]byte, len(s))
	top := -1

	for i := 0; i < len(s); i++ {
		if top >= 0 {
			if (stack[top] == '(' && s[i] == ')') || (stack[top] == '[' && s[i] == ']') ||
				(stack[top] == '{' && s[i] == '}') {
				top--
				continue
			}
		}
		top++
		stack[top] = s[i]
	}
	if top >= 0 {
		return false
	}
	return true
}

func main() {
	testcase := map[string]bool{
		"(":      false,
		"()":     true,
		"()[]{}": true,
		"(]":     false,
		"([)]":   false,
		"{[]}":   true,
	}
	for k, v := range testcase {
		m := isValid(k)
		if v != m {
			fmt.Println("test for", k, "fail, should be:", v, "but get", m)
		}
	}
}
