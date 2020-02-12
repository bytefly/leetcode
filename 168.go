package main

import (
	"fmt"
)

func convertToTitle(n int) string {
	var ans []byte
	if n < 1 {
		return ""
	}
	for n > 0 {
		if n%26 != 0 {
			ans = append([]byte{byte('A' + n%26 - 1)}, ans...)
		} else {
			ans = append([]byte{byte('Z')}, ans...)
			n -= 26
		}
		n /= 26
	}
	return string(ans)
}

func binary(n int) {
	for n > 0 {
		fmt.Println("-", n%16)
		n /= 16
	}
}

func main() {
	//binary(1053)
	fmt.Println(convertToTitle(1))
	fmt.Println(convertToTitle(26))
	fmt.Println(convertToTitle(27))
	fmt.Println(convertToTitle(28))
	fmt.Println(convertToTitle(52))
	fmt.Println(convertToTitle(701))
}
