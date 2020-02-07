package main

import (
	"fmt"
)

func mySqrt(x int) int {
	t := x / 2
	for t*t > x {
		t /= 2
	}

	stop := t * 2
	if t == stop {
		stop = x
	}

	ans := stop
	for t <= stop {
		m := t * t
		if m == x {
			ans = t
			break
		} else if m > x {
			ans = t - 1
			break
		}
		t++
	}

	return ans
}

func main() {
	fmt.Println(mySqrt(1))
	fmt.Println(mySqrt(2))
	fmt.Println(mySqrt(3))
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(5))
	fmt.Println(mySqrt(6))
	fmt.Println(mySqrt(7))
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(9))
	fmt.Println(mySqrt(10))
	fmt.Println(mySqrt(16))
}
