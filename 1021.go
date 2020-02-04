package main

import (
	"fmt"
)

func removeOuterParentheses(S string) string {
	var top string
	stack := make([]byte, 0)
	index, sIndex := 0, 0
	for i := 0; i < len(S); i++ {
		if S[i] == '(' { // handle '('
			stack = append(stack, S[i])
		} else { // handle ')'
			//pop left parenthese, no stack range check
			stack = stack[:len(stack)-1]
			if len(stack) == index {
				//drop outer parenthese
				top = S[sIndex+1 : i]
				stack = append(stack, top...)
				index += len(top)
				sIndex = i + 1
			}
		}
	}
	return string(stack)
}

func main() {
	fmt.Println(removeOuterParentheses("(()())(())(()(()))"))
	fmt.Println(removeOuterParentheses("(()(()))"))
	fmt.Println(removeOuterParentheses("(()())(())"))
	fmt.Println(removeOuterParentheses("()()"))
	fmt.Println(removeOuterParentheses("(())"))
}
