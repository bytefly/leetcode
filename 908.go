package main

import (
	"fmt"
	"math"
)

func smallestRangeI(A []int, K int) int {
	min, max := float64(A[0]), float64(A[0])
	for _, x := range A {
		min = math.Min(min, float64(x))
		max = math.Max(max, float64(x))
	}
	return int(math.Max(0, max-min-float64(2*K)))

}

func main() {
	fmt.Println(smallestRangeI([]int{1}, 0))
	fmt.Println(smallestRangeI([]int{0, 10}, 2))
	fmt.Println(smallestRangeI([]int{1, 3, 6}, 3))
	fmt.Println(smallestRangeI([]int{1, 3, 10}, 3))
	fmt.Println(smallestRangeI([]int{1, 3, 8, 10}, 3))
}
