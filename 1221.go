package main

import "fmt"

func balancedStringSplit(s string) int {
	var ans int
	var cnt [2]int

	for i := 0; i < len(s); i++ {
		if s[i] == 'L' {
			cnt[0]++
		} else {
			cnt[1]++
		}

		if cnt[0] == cnt[1] {
			ans++
			cnt[0], cnt[1] = 0, 0
		}
	}

	return ans
}

func main() {
	fmt.Println(balancedStringSplit("RLRRLLRLRL"))
	fmt.Println(balancedStringSplit("RLLLLRRRLR"))
	fmt.Println(balancedStringSplit("LLLLRRRR"))
	fmt.Println(balancedStringSplit("LR"))
	fmt.Println(balancedStringSplit("RRLRRLRLLLRL"))
}
