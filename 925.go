package main

import (
	"fmt"
)

func isLongPressedName(name string, typed string) bool {
	var i, j int
	for i < len(name) {
        if j >= len(typed) {
            return false
        }
		if j > 0 && i > 0 && name[i] != name[i-1] && typed[j-1] == typed[j] {
			j++
		} else if name[i] == typed[j] {
			i++
			j++
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isLongPressedName("alex", "aaleex"))
	fmt.Println(isLongPressedName("saeed", "ssaaedd"))
	fmt.Println(isLongPressedName("leelee", "lleeelee"))
	fmt.Println(isLongPressedName("laiden", "laiden"))
	fmt.Println(isLongPressedName("pyplrz", "ppyypllr"))
}
