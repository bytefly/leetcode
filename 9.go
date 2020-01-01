package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var num, prev int
	c := 10
	//convert form lower to higher
	for {
		//get current digit
		ch := x%c - prev
		ch /= c / 10
		//add to the num after multiple with 10
		num = num*10 + ch
		prev = x % c
		//stopped when get to the highest place
		if x/c == 0 {
			break
		}
		c *= 10
	}

	if x != num {
		return false
	}
	return true
}

func main() {
	testcase := map[int]bool{
		101:       true,
		1:         true,
		11:        true,
		121:       true,
		1001:      true,
		1221:      true,
		122353221: true,
		123:       false,
		-11:       false,
		-121:      false,
	}
	for k, v := range testcase {
		m := isPalindrome(k)
		if v != m {
			fmt.Println("test for '", k, "' fail, should be:", v, "but get", m)
		}
	}
}
