package main

import "fmt"

func gcdOfStrings(str1 string, str2 string) string {
	var ans string
	for i := 0; i < len(str1); i++ {
		match := true
		for j := i + 1; j < len(str1); j += i + 1 {
			if j+i+1 > len(str1) || str1[:i+1] != str1[j:j+i+1] {
				match = false
				break
			}
		}
		//str1 not match
		if !match {
			continue
		}

		if len(str2)%(i+1) == 0 {
			n := len(str2) / (i + 1)
			for j := 0; j < n; j += i + 1 {
				if str1[:i+1] != str2[j:j+i+1] {
					match = false
					break
				}
			}
		} else {
			continue
		}
		//str2 match, keep trying until find the longest one
		if match {
			ans = str1[:i+1]
		}
	}

	return ans
}

func main() {
	fmt.Println(gcdOfStrings("ABCABC", "ABCD"))
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))
	fmt.Println(gcdOfStrings("ABABAB", "ABAB"))
	fmt.Println(gcdOfStrings("ABC", "ABCABC"))
	fmt.Println(gcdOfStrings("ABAB", "ABABAB"))
	fmt.Println(gcdOfStrings("LEET", "CODE"))
	fmt.Println(gcdOfStrings("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"))
}
