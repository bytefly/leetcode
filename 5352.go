package main

import "fmt"

func generateTheString(n int) string {
	ans := make([]byte, n)
	if n%2 == 1 {
		for i := 0; i < n; i++ {
			ans[i] = 'a'
		}
	} else {
		for i := 0; i < n-1; i++ {
			ans[i] = 'a'
		}
		ans[n-1] = 'b'
	}

	return string(ans)
}

func main() {
	fmt.Println(generateTheString(4))
	fmt.Println(generateTheString(2))
	fmt.Println(generateTheString(7))
}
