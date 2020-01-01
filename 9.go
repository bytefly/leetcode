package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var revert int
	num := x
	//convert form lower to higher
	//stopped when get to the highest place
	for x > 0 {
		//add to the revert after multiple with 10
		revert = revert*10 + x%10
		x /= 10
	}

	if num != revert {
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
