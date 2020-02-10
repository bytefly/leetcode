package main

import (
	"fmt"
)

func countSegments(s string) int {
	var ans int
	findSpace := true
	for _, v := range s {
		if v == ' ' {
			if !findSpace {
				ans++
			}
			findSpace = true
		} else {
			findSpace = false
		}
	}
	if !findSpace {
		ans++
	}
	return ans
}

func main() {
	fmt.Println(countSegments(" a b c "))
	fmt.Println(countSegments("Hello"))
	fmt.Println(countSegments(" Hello "))
	fmt.Println(countSegments("Hello, my name is   John"))
	fmt.Println(countSegments("Hello, my name is   John "))
	fmt.Println(countSegments(" Hello, my name is   John "))
	fmt.Println(countSegments(" Hello,  my   name  is   John "))
}
