package main

import "fmt"

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	front, back := rec1, rec2
	up, down := rec1, rec2

	if rec2[0] < rec1[0] {
		front, back = back, front
	}
	if rec2[1] > rec1[1] {
		up, down = down, up
	}

	if back[0] < front[2] && up[1] < down[3] {
		return true
	}
	return false
}

func main() {
	fmt.Println(isRectangleOverlap([]int{0, 0, 2, 2}, []int{1, 1, 3, 3}))
	fmt.Println(isRectangleOverlap([]int{0, 0, 1, 1}, []int{1, 0, 2, 1}))
}
