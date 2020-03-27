package main

import "fmt"

func gcd(m, n int) int {
	if n == 0 {
		return m
	}
	return gcd(n, m%n)
}

func hasGroupsSizeX(deck []int) bool {
	var prevCnt, prevX int

	m := make(map[int]int)
	for _, num := range deck {
		m[num]++
	}

	for _, cnt := range m {
		if prevCnt == 0 {
			prevCnt = cnt
			if len(m) == 1 {
				prevX = cnt
			}
		} else {
			t := gcd(cnt, prevCnt)
			if prevX != 0 && t != prevX {
				return false
			}
			prevX = t
		}
	}

	if prevX == 1 {
		return false
	}

	return true
}

func main() {
	fmt.Println(hasGroupsSizeX([]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7}))
	fmt.Println(hasGroupsSizeX([]int{1, 2, 3, 4, 4, 3, 2, 1}))
	fmt.Println(hasGroupsSizeX([]int{1, 1, 1, 2, 2, 2, 3, 3}))
	fmt.Println(hasGroupsSizeX([]int{1}))
	fmt.Println(hasGroupsSizeX([]int{1, 1}))
	fmt.Println(hasGroupsSizeX([]int{1, 1, 2, 2, 2, 2}))
	fmt.Println(hasGroupsSizeX([]int{1, 1, 1, 1, 2, 2, 2, 2, 2, 2}))
}
