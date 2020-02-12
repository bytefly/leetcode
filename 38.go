package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	var num int
	ans := []byte{'1'}
	for i := 2; i <= n; i++ {
		size := len(ans)
		for j := 0; j < size; j++ {
			num++
			if j == size-1 || ans[j] != ans[j+1] {
				ans = append(ans, []byte(strconv.Itoa(num))...)
				ans = append(ans, ans[j])
				num = 0
			}
		}
		ans = ans[size:]
	}

	return string(ans)
}

func main() {
	fmt.Println(countAndSay(1))
	fmt.Println(countAndSay(2))
	fmt.Println(countAndSay(3))
	fmt.Println(countAndSay(4))
	fmt.Println(countAndSay(5))
	fmt.Println(countAndSay(6))
	fmt.Println(countAndSay(7))
	fmt.Println(countAndSay(8))
	fmt.Println(countAndSay(9))
	fmt.Println(countAndSay(30))
}
