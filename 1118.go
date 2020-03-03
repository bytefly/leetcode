package main

import "fmt"

func numberOfDays(Y int, M int) int {
	commonDays := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	isLeapYear := func(year int) bool {
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
			return true
		}
		return false
	}(Y)

	if !isLeapYear || M != 2 {
		return commonDays[M-1]
	}
	return 29
}

func main() {
	fmt.Println(numberOfDays(1992, 7))
	fmt.Println(numberOfDays(2000, 2))
	fmt.Println(numberOfDays(1900, 2))
}
