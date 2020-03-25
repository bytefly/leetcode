package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	var t [26]int

	m := make(map[[26]int][]string)
	for _, str := range strs {
		for _, b := range str {
			t[b-'a']++
		}

		m[t] = append(m[t], str)

		for i := 0; i < 26; i++ {
			t[i] = 0
		}
	}

	i, ans := 0, make([][]string, len(m))
	for _, v := range m {
		ans[i] = v
		i++
	}
	return ans
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
