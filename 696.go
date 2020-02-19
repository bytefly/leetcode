package main

import (
	"fmt"
)

func countBinarySubstrings(s string) int {
	var ans int
	var nums [2]int

	nums[s[0]-'0']++
	for i := 1; i < len(s); i++ {
		id := s[i] - '0'
		if s[i] != s[i-1] {
			nums[id] = 1
		} else {
			nums[id]++
		}

		if nums[id] <= nums[1-id] {
			ans++
		}
	}

	return ans
}

func main() {
	fmt.Println(countBinarySubstrings("00110011"))
	fmt.Println(countBinarySubstrings("10101"))
}
