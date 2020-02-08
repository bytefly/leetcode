package main

import (
	"fmt"
)

func getRow(rowIndex int) []int {
	numRow := rowIndex + 1
	ans := make([]int, numRow)
	if numRow < 1 {
		return []int{}
	}

	ans[0] = 1
	var t1, t2 int
	for i := 1; i < numRow; i++ {
		ans[i] = 1
		t1 = -1
		for j := 1; j < i; j++ {
			if t1 == -1 {
				t1 = 1
				t2 = ans[j]
				ans[j] += ans[j-1]
			} else {
				t1 = t2
				t2 = ans[j]
				ans[j] += t1
			}
		}
	}

	return ans
}

func main() {
	ans := getRow(9)
	fmt.Println(ans)
}
