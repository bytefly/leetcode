package main

import "fmt"

func removeVowels(S string) string {
	var ans []byte
	for _, b := range S {
		if b != 'a' && b != 'e' && b != 'i' && b != 'o' && b != 'u' {
			ans = append(ans, byte(b))
		}
	}

	return string(ans)
}

func main() {
	fmt.Println(removeVowels("leetcodeisacommunityforcoders"))
	fmt.Println(removeVowels("aeiou"))
}
