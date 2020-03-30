package main

import "fmt"

func lastRemaining(n int, m int) int {
	var t, ans int
	nums := make([]bool, n)

	if m == 1 {
		return n - 1
	}
	for {
		cnt := 0
		for i := 0; i < n; i++ {
			if !nums[i] {
				t = (t + 1) % m
				if t == 0 {
					nums[i] = true
				} else {
					ans = i
					cnt++
				}
			}
		}

		if cnt == 1 {
			break
		}
	}

	return ans
}

func main() {
	fmt.Println(lastRemaining(5, 1))
	fmt.Println(lastRemaining(5, 3))
	fmt.Println(lastRemaining(10, 17))
}
