package main

import (
	"fmt"
)

func addStrings(num1 string, num2 string) string {
	var ans string
	var carry byte
	i, j := len(num1)-1, len(num2)-1

	for i >= 0 || j >= 0 {
		t := carry
		if i >= 0 {
			t += num1[i] - '0'
			i--
		}
		if j >= 0 {
			t += num2[j] - '0'
			j--
		}

		if t > 9 {
			ans = string(t-10+'0') + ans
			carry = 1
		} else {
			ans = string(t+'0') + ans
			carry = 0
		}
	}

	if carry != 0 {
		ans = "1" + ans
	}
	return ans
}

func main() {
	fmt.Println(addStrings("0", "345"))
	fmt.Println(addStrings("99", "1"))
	fmt.Println(addStrings("123", "345"))
	fmt.Println(addStrings("999", "9999"))
}
