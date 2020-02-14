package main

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := make(map[byte]byte)
	tMap := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		v, ok := sMap[s[i]]
		if !ok {
			//value already be mapped
			v, ok = tMap[t[i]]
			if ok && v != s[i] {
				return false
			}
			sMap[s[i]] = t[i]
			tMap[t[i]] = s[i]
		} else {
			//map to different char
			if v != t[i] {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(isIsomorphic("ab", "ca"))
	fmt.Println(isIsomorphic("ab", "aa"))
	fmt.Println(isIsomorphic("aa", "ab"))
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("foo", "bar"))
	fmt.Println(isIsomorphic("paper", "title"))
}
