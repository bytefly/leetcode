package main

import (
	"fmt"
)

func isStrobogrammatic(num string) bool {
	left, right := 0, len(num)-1
	for left <= right {
		if num[left] == num[right] && (num[left] == '0' || num[left] == '1' || num[left] == '8') {
			left++
			right--
			continue
		} else if (num[left] == '6' && num[right] == '9') || (num[left] == '9' && num[right] == '6') {
			left++
			right--
			continue
		}
		return false
	}
	return true
}

func main() {
	fmt.Println(isStrobogrammatic("69"))
	fmt.Println(isStrobogrammatic("88"))
	fmt.Println(isStrobogrammatic("01"))
	fmt.Println(isStrobogrammatic("11"))
	fmt.Println(isStrobogrammatic("808"))
	fmt.Println(isStrobogrammatic("962"))
}
