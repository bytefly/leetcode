package main

import "fmt"

func removePalindromeSub(s string) int {
	var cnt [2]int
	var ans int

	isPalindrome := true
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			isPalindrome = false
			cnt[0]++
			cnt[1]++
		} else {
			cnt[s[left]-'a'] += 2
		}
		left++
		right--
	}
	if left == right {
		cnt[s[left]-'a']++
	}

	if len(s) > 0 && isPalindrome {
		return 1
	}

	for i := 0; i < 2; i++ {
		if cnt[i] > 0 {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(removePalindromeSub("ababa"))
	fmt.Println(removePalindromeSub("abb"))
	fmt.Println(removePalindromeSub("baabb"))
	fmt.Println(removePalindromeSub(""))
}
