package main

import (
	"fmt"
)

func thirdMax(nums []int) int {
	var first, second, third int
	var firstFound, secondFound, thirdFound bool

	for _, v := range nums {
		if !firstFound || v >= first {
			if !firstFound {
				first = v
				firstFound = true
			} else if v != first {
				if secondFound {
					third = second
					thirdFound = true
				}
				secondFound = true
				second = first
				first = v
			}
		} else if !secondFound || v >= second {
			if !secondFound {
				second = v
				secondFound = true
			} else if v != second {
				thirdFound = true
				third = second
				second = v
			}
		} else if !thirdFound || v > third {
			third = v
			if !thirdFound {
				thirdFound = true
			}
		}
	}

	if !thirdFound {
		third = first
	}
	return third
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
