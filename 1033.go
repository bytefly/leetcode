package main

import "fmt"

func numMovesStones(a int, b int, c int) []int {
	var min, max int

	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	if b > c {
		b, c = c, b
	}

	if b-a == 1 && c-b == 1 {
		min = 0
	} else if b-a <= 2 || c-b <= 2 {
		min = 1
	} else {
		min = 2
	}
	max = (c - b - 1) + (b - a - 1)

	return []int{min, max}
}

func main() {
	fmt.Println(numMovesStones(1, 2, 5))
	fmt.Println(numMovesStones(4, 3, 2))
	fmt.Println(numMovesStones(3, 5, 1))
}
