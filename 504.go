package main

import (
	"fmt"
)

func convertToBase7(num int) string {
	var n int
	abs := num
	if abs == 0 {
		return "0"
	}
	if abs < 0 {
		abs = -abs
	}

	var ans string
	for abs > 0 {
		abs, n = abs/7, abs%7
		ans = string('0'+n) + ans
	}

	if num < 0 {
		ans = "-" + ans
	}
	return ans
}

func main() {
	fmt.Println(convertToBase7(10000000))
	fmt.Println(convertToBase7(-10000000))
	fmt.Println(convertToBase7(100))
	fmt.Println(convertToBase7(-7))
	fmt.Println(convertToBase7(0))
	fmt.Println(convertToBase7(1))
}
