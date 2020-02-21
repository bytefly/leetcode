package main

import (
	"fmt"
)

func generatePossibleNextMoves(s string) []string {
	var ans []string
	bs := []byte(s)
	for i := 1; i < len(s); i++ {
		if s[i] == '+' && s[i] == s[i-1] {
			bs[i], bs[i-1] = '-', '-'
			ans = append(ans, string(bs))
			bs[i], bs[i-1] = '+', '+'
		}
	}

	return ans
}

func main() {
	fmt.Println(generatePossibleNextMoves("++++"))
	fmt.Println(generatePossibleNextMoves("+++++"))
}
