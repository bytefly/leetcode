package main

import (
	"fmt"
	"math"
)

func calculateTime(keyboard string, word string) int {
	var m [26]int
	var ans int

	for i, s := range keyboard {
		m[s-'a'] = i
	}

	ans = m[word[0]-'a']
	for i := 1; i < len(word); i++ {
		ans += int(math.Abs(float64(m[word[i]-'a'] - m[word[i-1]-'a'])))
	}

	return ans
}

func main() {
	fmt.Println(calculateTime("abcdefghijklmnopqrstuvwxyz", "cba"))
	fmt.Println(calculateTime("pqrstuvwxyzabcdefghijklmno", "leetcode"))
}
