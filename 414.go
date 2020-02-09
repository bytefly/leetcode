package main

import (
	"fmt"
	"math"
)

func thirdMax(nums []int) int {
	first, second, third := math.Inf(-1), math.Inf(-1), math.Inf(-1)

	for _, i := range nums {
		v := float64(i)
		if v >= first {
			if v != first {
				third = second
				second = first
				first = v
			}
		} else if v >= second {
			if v != second {
				third = second
				second = v
			}
		} else if v > third {
			third = v
		}
	}

	if math.IsInf(third, -1) {
		third = first
	}
	return int(third)
}

func main() {
	fmt.Println(thirdMax([]int{3, 2, 1}))
	fmt.Println(thirdMax([]int{1, 2}))
	fmt.Println(thirdMax([]int{1}))
	fmt.Println(thirdMax([]int{1, 1}))
	fmt.Println(thirdMax([]int{1, 1, 1}))
	fmt.Println(thirdMax([]int{5, 2, 2}))
	fmt.Println(thirdMax([]int{2, 2, 3, 1}))
	fmt.Println(thirdMax([]int{1, 2, 2, 5, 3, 5}))
	fmt.Println(thirdMax([]int{1, 2, -2147483648}))
}
