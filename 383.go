package main

import (
	"fmt"
)

func canConstruct(ransomNote string, magazine string) bool {
	noteMap := make(map[byte]int)
	letterMap := make(map[byte]int)
	for i := 0; i < len(ransomNote); i++ {
		v, ok := noteMap[ransomNote[i]]
		if !ok {
			noteMap[ransomNote[i]] = 1
		} else {
			noteMap[ransomNote[i]] = v + 1
		}
	}
	for i := 0; i < len(magazine); i++ {
		v, ok := letterMap[magazine[i]]
		if !ok {
			letterMap[magazine[i]] = 1
		} else {
			letterMap[magazine[i]] = v + 1
		}
	}

	for k, v := range noteMap {
		if val, ok := letterMap[k]; !ok || val < v {
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
