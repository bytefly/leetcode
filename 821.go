package main

import "fmt"

func shortestToChar(S string, C byte) []int {
	ans := make([]int, len(S))
	for i := 0; i < len(S); i++ {
		if S[i] == C {
			ans[i] = 0
			for j := i - 1; j >= 0 && S[j] != C; j-- {
				if ans[j] == 0 || ans[j] > i-j {
					ans[j] = i - j
				}
			}
			for j := i + 1; j < len(S) && S[j] != C; j++ {
				if ans[j] == 0 || j-i < ans[j] {
					ans[j] = j - i
				}
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(shortestToChar("loveleetcode", 'e'))
}
