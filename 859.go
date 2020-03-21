package main

import "fmt"

func buddyStrings(A string, B string) bool {
	var diff [2]int
	var m [26]int
	var j int

	if len(A) != len(B) {
		return false
	}
	if A == B {
		for i := 0; i < len(A); i++ {
			//it's ok if it has same char
			if m[A[i]-'a'] == 1 {
				return true
			}
			m[A[i]-'a']++
		}
		return false
	}

	for i := 0; i < len(A); i++ {
		if A[i] != B[i] {
			//it's not ok if it has more than 2 chars in different position
			if j == 2 {
				return false
			}
			diff[j] = i
			j++
		}
	}

	//different chars in two position
	if A[diff[0]] != B[diff[1]] || A[diff[1]] != B[diff[0]] {
		return false
	}
	return true
}

func main() {
	fmt.Println(buddyStrings("abab", "abab"))
	fmt.Println(buddyStrings("abcd", "abcd"))
	fmt.Println(buddyStrings("ab", "ba"))
	fmt.Println(buddyStrings("ab", "ab"))
	fmt.Println(buddyStrings("aa", "aa"))
	fmt.Println(buddyStrings("aaaaaaabc", "aaaaaaacb"))
	fmt.Println(buddyStrings("", "aa"))
	fmt.Println(buddyStrings("aaaaaaabc", "aaaaaaacd"))
	fmt.Println(buddyStrings("aaaaaaabcd", "aaaaaaadbc"))
}
