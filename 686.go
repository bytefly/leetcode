package main

import (
	"fmt"
	"strings"
)

func repeatedStringMatch(A string, B string) int {
	var matchCnt int

	firstPos := strings.Index(B, A)
	if firstPos < 0 {
		if strings.Index(A, B) < 0 {
			// B is in the middle of A+A
			if strings.Index(A+A, B) >= 0 {
				return 2
			}
			return -1
		} else { // B is a part of A
			return 1
		}
	}

	lastPos := firstPos
	matchCnt++
	for {
		pos := strings.Index(B[lastPos+len(A):], A)
		if pos > 0 && pos != lastPos+len(A) {
			//has gap between As in B
			return -1
		}
		if pos == -1 {
			break
		}
		lastPos += len(A) + pos
		matchCnt++
	}

	if firstPos > 0 {
		if strings.Index(A, B[:firstPos]) > 0 {
			matchCnt++
		} else { //not end with B[:firstPos]
			return -1
		}
	}
	if lastPos < len(B)-len(A) {
		if strings.Index(A, B[lastPos+len(A):]) == 0 {
			matchCnt++
		} else { //not begin with B[lastPos+len(A):]
			return -1
		}
	}
	return matchCnt
}

func main() {
	fmt.Println(repeatedStringMatch("aaaaaaaaaaaaaaaaaaaaaab", "ba"))
	fmt.Println(repeatedStringMatch("abcd", "ab"))
	fmt.Println(repeatedStringMatch("abcd", "cd"))
	fmt.Println(repeatedStringMatch("abcd", "cdabcdab"))
	fmt.Println(repeatedStringMatch("abcd", "cdabcd"))
	fmt.Println(repeatedStringMatch("abcd", "abcdab"))
	fmt.Println(repeatedStringMatch("abcd", "cdabcdabcdab"))
	fmt.Println(repeatedStringMatch("abcd", "cdabcdaabcdab"))
	fmt.Println(repeatedStringMatch("abcd", "ababcdab"))
	fmt.Println(repeatedStringMatch("abcd", "cdabcdcd"))
}
