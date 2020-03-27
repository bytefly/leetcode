package main

import "fmt"

func gcd(m, n int) int {
	if n == 0 {
		return m
	}
	return gcd(n, m%n)
}

func hasGroupsSizeX(deck []int) bool {
	m := make(map[int]int)
	for _, num := range deck {
		m[num]++
	}

	x := m[deck[0]]
	for _, cnt := range m {
		if cnt != x {
			x = gcd(cnt, x)
		}
		if x == 1 {
			return false
		}
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
