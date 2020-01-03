package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	t := strs[0]

	for i := 1; i <= len(t); i++ {
		var j int
		for j = 1; j < len(strs); j++ {
			if len(strs[j]) < i {
				break
			}
			if strings.Compare(t[0:i], strs[j][0:i]) != 0 {
				break
			}
		}

		if i == 1 && j == 1 {
			return ""
		} else if j != len(strs) {
			return t[0 : i-1]
		}
	}
	return t
}

func main() {
	s := []string{"flower", "flow", "flight"}
	//s := []string{"dog", "racecar", "car"}
	v := longestCommonPrefix(s)
	fmt.Println(v)
}
