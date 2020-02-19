package main

import (
	"fmt"
)

func isOneBitCharacter(bits []int) bool {
	var i int
	length := len(bits)
	for i < length {
		if i == length-1 {
			return true
		}
		if bits[i] == 0 {
			i++
		} else {
			i += 2
		}
	}
	return false
}

func main() {
	fmt.Println(isOneBitCharacter([]int{0}))
	fmt.Println(isOneBitCharacter([]int{1, 0}))
	fmt.Println(isOneBitCharacter([]int{1, 0, 0}))
	fmt.Println(isOneBitCharacter([]int{1, 1, 1, 0}))
}
