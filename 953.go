package main

import "fmt"

func isAlienSorted(words []string, order string) bool {
	var m [26]int

	for i := 0; i < 26; i++ {
		m[order[i]-'a'] = i
	}

	for j := 1; j < len(words); j++ {
		nextLoop := false
		for i := 0; i < len(words[j-1]) && i < len(words[j]); i++ {
			s, t := words[j-1][i]-'a', words[j][i]-'a'
			if s != t {
				if m[s] > m[t] {
					return false
				}
				nextLoop = true
				break
			}
		}

		if !nextLoop && len(words[j]) < len(words[j-1]) {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"))
	fmt.Println(isAlienSorted([]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz"))
	fmt.Println(isAlienSorted([]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(isAlienSorted([]string{"kuvp", "q"}, "ngxlkthsjuoqcpavbfdermiywz"))
}
