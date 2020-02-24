package main

import "fmt"

func isNoZero(n int) bool {
	for n > 0 {
		if n%10 == 0 {
			return false
		}
		n /= 10
	}
	return true
}

func getNoZeroIntegers(n int) []int {
	num := n / 2
	for {
		if isNoZero(num) && isNoZero(n-num) {
			return []int{num, n - num}
		}
		num -= 1
	}
}

func main() {
	fmt.Println(getNoZeroIntegers(2))
	fmt.Println(getNoZeroIntegers(3))
	fmt.Println(getNoZeroIntegers(4))
	fmt.Println(getNoZeroIntegers(11))
	fmt.Println(getNoZeroIntegers(10000))
	fmt.Println(getNoZeroIntegers(69))
	fmt.Println(getNoZeroIntegers(1010))
}
