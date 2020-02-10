package main

import (
	"fmt"
)

func findComplement(num int) int {
	var pos uint
	if num == 0 {
		return 1
	}
	for pos = 0; pos < 32; pos++ {
		if num <= (1 << pos) {
			if num != (1 << pos) {
				pos -= 1
			}
			break
		}
	}

	return (1<<(pos+1) - 1) & ^num
}

func main() {
	fmt.Println(findComplement(2048))
	fmt.Println(findComplement(8))
	fmt.Println(findComplement(5))
	fmt.Println(findComplement(1))
	fmt.Println(findComplement(0))
}
