package main

import "fmt"

func checkStraightLine(coordinates [][]int) bool {
	for i := 2; i < len(coordinates); i++ {
		if (coordinates[i][1]-coordinates[i-1][1])*(coordinates[i-1][0]-coordinates[i-2][0]) != (coordinates[i-1][1]-coordinates[i-2][1])*(coordinates[i][0]-coordinates[i-1][0]) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(checkStraightLine([][]int{
		[]int{1, 2}, []int{2, 3}, []int{3, 4}, []int{4, 5}, []int{5, 6}, []int{6, 7},
	}))
	fmt.Println(checkStraightLine([][]int{
		[]int{1, 1}, []int{2, 2}, []int{3, 4}, []int{4, 5}, []int{5, 6}, []int{7, 7},
	}))
}
