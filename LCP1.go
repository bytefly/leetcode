package main

import "fmt"

func game(guess []int, answer []int) int {
	ans := 3
	for i := 0; i < 3; i++ {
		if guess[i] != answer[i] {
			ans--
		}
	}
	return ans
}

func main() {
	fmt.Println(game([]int{1, 2, 3}, []int{1, 2, 3}))
	fmt.Println(game([]int{2, 2, 3}, []int{3, 2, 1}))
}
