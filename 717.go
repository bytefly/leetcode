package main

import (
	"fmt"
)

func isOneBitCharacter(bits []int) bool {
	var i int
	length := len(bits)
	for i < length-1 {
		i += bits[i] + 1
	}
	return i == length-1
}

func main() {
	fmt.Println(isOneBitCharacter([]int{0}))
	fmt.Println(isOneBitCharacter([]int{1, 0}))
	fmt.Println(isOneBitCharacter([]int{1, 0, 0}))
	fmt.Println(isOneBitCharacter([]int{1, 1, 1, 0}))
}
