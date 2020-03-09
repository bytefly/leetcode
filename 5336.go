package main

import "fmt"

func sortString(s string) string {
	var ans []byte
	var m [26][]int
	var empty int

	for i := 0; i < len(s); i++ {
		m[s[i]-'a'] = append(m[s[i]-'a'], i)
	}

	for empty < 26 {
		for i := 0; i < 26; i++ {
			if len(m[i]) > 0 {
				ans = append(ans, byte('a'+i))
				m[i] = m[i][1:]
			}
		}
		empty = 0
		for i := 25; i >= 0; i-- {
			if len(m[i]) > 0 {
				ans = append(ans, byte('a'+i))
				m[i] = m[i][1:]
			} else {
				empty++
			}
		}
	}

	return string(ans)
}

func main() {
	fmt.Println(sortString("aaaabbbbcccc"))
	fmt.Println(sortString("rat"))
	fmt.Println(sortString("leetcode"))
	fmt.Println(sortString("ggggggg"))
	fmt.Println(sortString("spo"))
}
