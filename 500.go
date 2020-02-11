package main

import (
	"fmt"
	"strings"
)

func findWords(words []string) []string {
	var ans []string
	var counter [3]int
	rowMap := [26]int{1, 2, 2, 1, 0, 1, 1, 1, 0, 1, 1, 1, 2, 2, 0, 0, 0, 0, 1, 0, 0, 2, 0, 2, 0, 2}

	for m, i := range words {
		i = strings.ToUpper(i)
		for _, j := range i {
			counter[rowMap[j-'A']]++
		}
		for j := 0; j < 3; j++ {
			if counter[j] == len(i) {
				counter[j] = 0
				ans = append(ans, words[m])
				break
			}
			counter[j] = 0
		}
	}

	return ans
}

func main() {
	fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
}
