package main

import "fmt"

func canThreePartsEqualSum(A []int) bool {
	var partSum, cnt, t int

	for _, num := range A {
		partSum += num
	}
	if partSum%3 != 0 {
		return false
	}
	partSum /= 3

	for _, num := range A {
		t += num
		if t == partSum {
			t = 0
			cnt++
		}
	}
	if cnt < 3 {
		return false
	}
	return true
}

func main() {
	fmt.Println(canThreePartsEqualSum([]int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1}))
	fmt.Println(canThreePartsEqualSum([]int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1}))
	fmt.Println(canThreePartsEqualSum([]int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4}))
}
