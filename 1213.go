package main

import "fmt"

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
	var cnt [2001]int
	var ans []int

	for _, num := range arr1 {
		cnt[num]++
	}
	for _, num := range arr2 {
		cnt[num]++
	}
	for _, num := range arr3 {
		if cnt[num] >= 2 {
			ans = append(ans, num)
		}
	}

	return ans
}

func main() {
	fmt.Println(arraysIntersection([]int{1, 2, 3, 4, 5}, []int{1, 2, 5, 7, 9}, []int{1, 3, 4, 5, 8}))
}
