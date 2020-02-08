package main

import (
	"fmt"
)

func canConstruct(ransomNote string, magazine string) bool {
	var noteMap [26]int
	var letterMap [26]int
	for i := 0; i < len(ransomNote); i++ {
			noteMap[ransomNote[i]-'a']++
	}
	for i := 0; i < len(magazine); i++ {
			letterMap[magazine[i]-'a']++
	}

	for k, v := range noteMap {
        if v > letterMap[k] {
            return false
        }
	}
	return true
}

func main() {
	fmt.Println(canConstruct("a", "b"))
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("aa", "aba"))
}
