package main

import (
	"fmt"
	"strconv"
	"strings"
)

func daysBetweenDates(date1 string, date2 string) int {
	var ans int
	commonDays := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	meta1 := strings.Split(date1, "-")
	meta2 := strings.Split(date2, "-")
	start, end := meta1, meta2

	for i := 0; i < 3; i++ {
		if start[i] < end[i] {
			break
		}
		if start[i] > end[i] {
			start, end = end, start
			break
		}
	}
	//check if year is leap if it's divisable by 4 not 100 or by 400
	isLeapYear := func(year int) bool {
		if year%4 != 0 {
			return false
		}
		if year%100 != 0 {
			return true
		}
		if year%400 != 0 {
			return false
		}
		return true
	}

	startYear, _ := strconv.Atoi(start[0])
	endYear, _ := strconv.Atoi(end[0])
	startMon, _ := strconv.Atoi(start[1])
	endMon, _ := strconv.Atoi(end[1])
	startDay, _ := strconv.Atoi(start[2])
	endDay, _ := strconv.Atoi(end[2])

	//compute middle complete years' days
	for i := startYear + 1; i <= endYear-1; i++ {
		ans += 365
		if isLeapYear(i) {
			ans++
		}
	}

	if endYear != startYear {
		//compute start month last days
		var startNextDays, endNextDays int
		for i := startMon; i <= 12; i++ {
			startNextDays += commonDays[i-1]
			if i == 2 && isLeapYear(startYear) {
				startNextDays++
			}
		}
		startNextDays -= startDay
		//compute end month start days
		for i := 1; i < endMon; i++ {
			endNextDays += commonDays[i-1]
			if i == 2 && isLeapYear(endYear) {
				endNextDays++
			}
		}
		endNextDays += endDay

		ans += startNextDays + endNextDays
	} else {
		for i := startMon; i < endMon; i++ {
			ans += commonDays[i-1]
			if i == 2 && isLeapYear(startYear) {
				ans++
			}
		}
		ans -= startDay
		ans += endDay
	}

	return ans
}

func main() {
	fmt.Println(daysBetweenDates("1971-06-29", "2010-09-23"))
	fmt.Println(daysBetweenDates("2019-06-29", "2019-06-30"))
	fmt.Println(daysBetweenDates("2020-01-15", "2019-12-31"))
}
