package main

import (
	"fmt"
)

func judgeCircle(moves string) bool {
	var x, y int
	for _, v := range moves {
		switch v {
		case 'U':
			y++
		case 'D':
			y--
		case 'L':
			x--
		case 'R':
			x++
		}
	}

	if x == 0 && y == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(judgeCircle("UD"))
	fmt.Println(judgeCircle("LL"))
}
