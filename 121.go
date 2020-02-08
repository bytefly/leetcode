package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	max := 0
	for i := 0; i < len(prices); i++ {
		if minPrice >= prices[i] {
			minPrice = prices[i]
		} else {
			t := prices[i] - minPrice
			if t > max {
				max = t
			}
		}
	}
	return max
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 2, 5, 3, 6, 4, 1}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
