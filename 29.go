package main

import (
	"fmt"
)

const (
	intMax = int(^uint32(0) >> 1)
	intMin = -intMax - 1
)

func divide(dividend int, divisor int) int {
	var t, quot, n int
	var negative bool

	if divisor == -1 && dividend == intMin {
		return intMax
	}
	if dividend < 0 && divisor > 0 {
		negative = true
		dividend = -dividend
	} else if dividend >= 0 && divisor < 0 {
		negative = true
		divisor = -divisor
	} else if dividend < 0 && divisor < 0 {
		divisor = -divisor
		dividend = -dividend
	}
	for {
		if divisor<<n < dividend {
			n++
		} else {
			break
		}
	}

	if n == 0 {
		t = dividend
	} else {
		t = dividend - divisor<<(n-1)
		quot = 1 << (n - 1)
	}
	//fmt.Println(dividend, divisor, quot, t)
	for t >= divisor {
		t -= divisor
		quot++
	}
	if negative {
		return -quot
	}
	return quot
}

func main() {
	m := []int{15, 2, 7,
		2, 5, 0,
		8, 2, 4,
		-2, 5, 0,
		-8, 2, -4,
		8, -2, -4,
		-15, 2, -7,
		15, -2, -7,
		-15, -2, 7,
		10, 3, 3,
		7, -3, -2,
		1, 1, 1,
		intMin, -1, intMax,
	}
	for i := 0; i < len(m); {
		v := divide(m[i], m[i+1])
		if v != m[i+2] {
			fmt.Println("test fail, should be:", m[i+2], "but get", v)
		}
		i += 3
	}
}
