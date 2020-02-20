package main

import (
	"fmt"
)

func numJewelsInStones(J string, S string) int {
	var ans int
	var jewels [52]bool

	for _, v := range J {
		if v > 'Z' {
			jewels[v-'a'] = true
		} else {
			jewels[v-'A'+26] = true
		}
	}

	for _, v := range S {
		if v > 'Z' {
			if jewels[v-'a'] {
				ans++
			}
		} else {
			if jewels[v-'A'+26] {
				ans++
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
	fmt.Println(numJewelsInStones("z", "ZZ"))
	fmt.Println(numJewelsInStones("azAZ", "azAZZ"))
}
