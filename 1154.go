package main

import (
	"fmt"
	"strconv"
	"strings"
)

func dayOfYear(date string) int {
	commonDays := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	nums := strings.Split(date, "-")
	year, _ := strconv.Atoi(nums[0])
	month, _ := strconv.Atoi(nums[1])
	day, _ := strconv.Atoi(nums[2])
	ans := day

	isLeapYear := func() bool {
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
			return true
		}
		return false
	}()

	for i := 1; i < month; i++ {
		ans += commonDays[i-1]
	}
	if isLeapYear && month > 2 {
		ans++
	}

	return ans
}

func main() {
	fmt.Println(dayOfYear("2019-01-09"))
	fmt.Println(dayOfYear("2019-02-10"))
	fmt.Println(dayOfYear("2003-03-01"))
	fmt.Println(dayOfYear("2004-03-01"))
}
