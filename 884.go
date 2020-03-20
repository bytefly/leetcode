package main

import "fmt"

func uncommonFromSentences(A string, B string) []string {
	var ans []string
	m := make(map[string]int)

	start := -1
	for i := 0; i < len(A); i++ {
		if A[i] != ' ' && start < 0 {
			start = i
		} else if A[i] == ' ' && start >= 0 {
			m[A[start:i]]++
			start = -1
		}
	}
	if A[len(A)-1] != ' ' && start >= 0 {
		m[A[start:]]++
	}

	start = -1
	for i := 0; i < len(B); i++ {
		if B[i] != ' ' && start < 0 {
			start = i
		} else if B[i] == ' ' && start >= 0 {
			m[B[start:i]]++
			start = -1
		}
	}
	if B[len(B)-1] != ' ' && start >= 0 {
		m[B[start:]]++
	}

	for k, v := range m {
		if v == 1 {
			ans = append(ans, k)
		}
	}
	return ans
}

func main() {
	fmt.Println(uncommonFromSentences("this apple is sweet", "this apple is sour"))
	fmt.Println(uncommonFromSentences("apple apple", "banana"))
	fmt.Println(uncommonFromSentences("fd kss fd", "fd fd kss"))
}
