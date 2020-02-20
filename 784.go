package main

import (
	"fmt"
)

func letterCasePermutation(S string) []string {
	var ans []string
	var cnt int

	for i := 0; i < len(S); i++ {
		if S[i] > '9' || S[i] < '0' {
			//last byte is a char
			if i == len(S)-1 {
				ans = append(ans, S[:i+1])
				ans = append(ans, S[:i]+string(S[i]^32))
			} else {
				m := letterCasePermutation(S[i+1:])
				for _, s := range m {
					ans = append(ans, S[:i+1]+s)
					ans = append(ans, S[:i]+string(S[i]^32)+s)
				}
			}
			break
		} else {
			cnt++
		}
	}
	//all numbers in S
	if cnt == len(S) {
		ans = append(ans, S)
	}

	return ans
}

func main() {
	fmt.Println(letterCasePermutation("a"))
	fmt.Println(letterCasePermutation("abc"))
	fmt.Println(letterCasePermutation("a1b2"))
	fmt.Println(letterCasePermutation("a12b"))
	fmt.Println(letterCasePermutation("3z4"))
	fmt.Println(letterCasePermutation("12345"))
}
