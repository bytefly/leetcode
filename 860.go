package main

import (
	"fmt"
)

func lemonadeChange(bills []int) bool {
	var income [5]int
	for _, bill := range bills {
		if bill != 5 {
			change := bill - 5
			for i := 4; i > 0; i-- {
				m := change / i / 5
				if m > 0 && m <= income[i] {
					income[i] -= m
					change -= m * 5 * i
				}
			}
			if change > 0 {
				return false
			}
		}

		income[bill/5]++
	}

	return true
}

func main() {
	fmt.Println(lemonadeChange([]int{5, 5, 5, 10, 20}))
	fmt.Println(lemonadeChange([]int{5, 5, 10}))
	fmt.Println(lemonadeChange([]int{10, 10}))
	fmt.Println(lemonadeChange([]int{5, 5, 10, 10, 20}))
}
