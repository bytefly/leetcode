package main

import (
	"fmt"
)

func canAttendMeetings(intervals [][]int) bool {
	for i := 0; i < len(intervals); i++ {
		for j := i + 1; j < len(intervals); j++ {
			if !(intervals[j][1] <= intervals[i][0] || intervals[j][0] >= intervals[i][1]) {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(canAttendMeetings([][]int{[]int{0, 30}, []int{5, 10}, []int{15, 20}}))
	fmt.Println(canAttendMeetings([][]int{[]int{7, 10}, []int{2, 4}}))
	fmt.Println(canAttendMeetings([][]int{[]int{13, 15}, []int{1, 13}}))
}
