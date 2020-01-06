package main

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}

	ans := make([]string, 0)
	for c := 0; c < n; c++ {
		lans := generateParenthesis(c)
		for _, left := range lans {
			rans := generateParenthesis(n - 1 - c)
			for _, right := range rans {
				ans = append(ans, "("+left+")"+right)
			}
		}
	}

	return ans
}

func main() {
	s := generateParenthesis(3)
	for _, t := range s {
		fmt.Println(t)
	}
}
