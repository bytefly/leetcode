package main

import (
	"fmt"
)

func toHex(num int) string {
	var ans string
	var numEx uint32

	if num < 0 {
		numEx = uint32(^(-num) + 1)
	} else {
		numEx = uint32(num)
	}

	for numEx >= 0 {
		quote := numEx % 16
		if quote < 10 {
			ans = string(quote+'0') + ans
		} else {
			ans = string(quote-10+'a') + ans
		}
		numEx /= 16

		if numEx == 0 {
			break
		}
	}

	return ans
}

func main() {
	fmt.Println(toHex(0))
	fmt.Println(toHex(16))
	fmt.Println(toHex(26))
	fmt.Println(toHex(2002))
	fmt.Println(toHex(-5))
	fmt.Println(toHex(-1))
}
