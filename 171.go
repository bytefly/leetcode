package main

import (
	"fmt"
)

func titleToNumber(s string) int {
	var ans int
	for i := 0; i < len(s); i++ {
		ans *= 26
		ans += int(s[i] - 'A') + 1
	}
	return ans
}

func main() {
	fmt.Println(titleToNumber("A"))
	fmt.Println(titleToNumber("AB"))
	fmt.Println(titleToNumber("AZ"))
	fmt.Println(titleToNumber("ZY"))
}
