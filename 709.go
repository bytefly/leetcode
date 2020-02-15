package main

import (
	"fmt"
)

func toLowerCase(str string) string {
	ans := []byte(str)
	for i := 0; i < len(ans); i++ {
		if ans[i] >= 'A' && ans[i] <= 'Z' {
			ans[i] -= 'A'
			ans[i] += 'a'
		}
	}
	return string(ans)
}

func main() {
	fmt.Println(toLowerCase("Hello"))
	fmt.Println(toLowerCase("here"))
	fmt.Println(toLowerCase("LOVELY"))
}
