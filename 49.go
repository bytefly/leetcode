package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	var ans [][]string
	var t [26]int

	m := make(map[[26]int][]string)
	for _, str := range strs {
		for _, b := range str {
			t[b-'a']++
		}

		_, ok := m[t]
		if ok {
			m[t] = append(m[t], str)
		} else {
			m[t] = []string{str}
		}

		for i := 0; i < 26; i++ {
			t[i] = 0
		}
	}

	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
