package main

import (
	"fmt"
)

func maximum69Number(num int) int {
	var cnt, m int
	n, max := num, num

	for n > 0 {
		if n%10 == 6 {
			if cnt > 0 {
				prev := 1
				for i := 0; i < cnt; i++ {
					prev *= 10
				}
				m = (n+3)*prev + num%prev
			} else {
				m = n + 3
			}

			if m > max {
				max = m
			}
		}

		n /= 10
		cnt++
	}
	return max
}

func main() {
	fmt.Println(maximum69Number(6))
	fmt.Println(maximum69Number(9))
	fmt.Println(maximum69Number(9669))
	fmt.Println(maximum69Number(9996))
	fmt.Println(maximum69Number(9999))
}
