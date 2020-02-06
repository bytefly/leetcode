package main

import (
	"fmt"
)

func isPerfectSquare(num int) bool {
	t := num / 2
	for t*t > num {
		t /= 2
	}

	stop := t * 2
	if t == stop {
		stop = num
	}
	for t <= stop {
		if t*t == num {
			return true
		} else {
			t++
		}
	}
	return false
}

func main() {
	fmt.Println(isPerfectSquare(1))
	fmt.Println(isPerfectSquare(2))
	fmt.Println(isPerfectSquare(3))
	fmt.Println(isPerfectSquare(4))
	fmt.Println(isPerfectSquare(9))
	fmt.Println(isPerfectSquare(16))
	fmt.Println(isPerfectSquare(14))
}
