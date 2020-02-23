package main

import (
	"fmt"
)

func diStringMatch(S string) []int {
	var iPos, ans []int
	var dLen, next int

	//save Is' positions
	for i := 0; i < len(S); i++ {
		if S[i] == 'I' {
			iPos = append(iPos, i)
		}
	}

	//fill Is' numbers first
	for i := 0; i <= len(iPos); i++ {
		ans = append(ans, i)
	}
	next = len(iPos) + 1

	//fill numbers of Ds which after each I
	for i := len(iPos) - 1; i >= 0; i-- {
		if i == len(iPos)-1 {
			dLen = len(S) - iPos[i] - 1
		} else {
			dLen = iPos[i+1] - iPos[i] - 1
		}

		for j := 0; j < dLen; j++ {
			ans = append(ans[:i+1], append([]int{next}, ans[i+1:]...)...)
			next++
		}
	}

	//fill numbers of Ds at the beginning
	if len(iPos) == 0 || iPos[0] > 0 {
		size := len(S)
		if len(iPos) != 0 {
			size = iPos[0]
		}
		for j := 0; j < size; j++ {
			ans = append([]int{next}, ans[:]...)
			next++
		}
	}

	return ans
}

func main() {
	fmt.Println(diStringMatch("IDID"))
	fmt.Println(diStringMatch("III"))
	fmt.Println(diStringMatch("DDI"))
	fmt.Println(diStringMatch("IDD"))
	fmt.Println(diStringMatch("DDIDD"))
	fmt.Println(diStringMatch("D"))
	fmt.Println(diStringMatch("DDD"))
}
