package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	var ans []byte
	var t, carry byte

	if len(a) < len(b) {
		a, b = b, a
	}

	offs := len(a) - len(b)
	for i := len(a) - 1; i >= 0; i-- {
		if i-offs < 0 {
			t = a[i] - '0' + carry
		} else {
			t = a[i] - '0' + b[i-offs] - '0' + carry
		}
		if t >= 2 {
			ans = append([]byte{'0' + t - 2}, ans...)
			carry = 1
		} else {
			ans = append([]byte{'0' + t}, ans...)
			carry = 0
		}
	}

	if carry > 0 {
		ans = append([]byte{'1'}, ans...)
	}

	return string(ans)
}

func main() {
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1010", "1011"))
	fmt.Println(addBinary("1111", "1111"))
	fmt.Println(addBinary("1111", "1111111111111111111111011"))
}
