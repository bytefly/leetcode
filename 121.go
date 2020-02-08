package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	max := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if max < prices[j]-prices[i] {
				max = prices[j] - prices[i]
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
