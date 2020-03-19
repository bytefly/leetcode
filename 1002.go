package main

import "fmt"

func commonChars(A []string) []string {
	var ans []string
	var m [2][26]int

	for _, b := range A[0] {
		m[0][b-'a']++
	}
	for i := 1; i < len(A); i++ {
		for _, b := range A[i] {
			m[1][b-'a']++
		}
		for j := 0; j < 26; j++ {
			if m[0][j] > m[1][j] {
				m[0][j] = m[1][j]
			}
			m[1][j] = 0
		}
	}

	for j := 0; j < 26; j++ {
		for i := 0; i < m[0][j]; i++ {
			ans = append(ans, string('a'+j))
		}
	}
	return ans
}

func main() {
	fmt.Println(commonChars([]string{"bella", "label", "roller"}))
	fmt.Println(commonChars([]string{"cool", "lock", "cook"}))
	fmt.Println(commonChars([]string{"cool"}))
}
