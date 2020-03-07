package main

import (
	"fmt"
)

func numberOfLines(widths []int, S string) []int {
	var sum, row int
	for i := 0; i < len(S); i++ {
		sum += widths[S[i]-'a']
		if sum >= 100 {
			if sum > 100 {
				i--
			}
			sum = 0
			row++
		}
	}
	if sum > 0 {
		row++
	}

	return []int{row, sum}
}

func main() {
	fmt.Println(numberOfLines([]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(numberOfLines([]int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "bbbcccdddaaa"))
}
