package main

import "fmt"

func missingNumber(arr []int) int {
	var ans int
	for i := 1; i < len(arr)-1; i++ {
		m, n := arr[i+1]-arr[i], arr[i]-arr[i-1]
		if m != n {
			ans = arr[i] + (m - n)
			break
		}
	}
	return ans
}

func main() {
	fmt.Println(missingNumber([]int{5, 7, 11, 13}))
	fmt.Println(missingNumber([]int{5, 9, 11, 13}))
	fmt.Println(missingNumber([]int{15, 13, 12}))
	fmt.Println(missingNumber([]int{15, 14, 12}))
}
