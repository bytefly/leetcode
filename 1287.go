package main

import "fmt"

func findSpecialInteger(arr []int) int {
	size := len(arr) / 4
	for i := 1; i <= 2; i++ {
		id := size * i
		left := searchLeft(arr, arr[id])
		right := searchRight(arr, arr[id])
		if right-left >= size {
			return arr[id]
		}
	}

	return arr[size*3]
}

func searchLeft(arr []int, v int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := (left + right) / 2
		if arr[mid] >= v {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func searchRight(arr []int, v int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := (left + right) / 2
		if arr[mid] > v {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left - 1
}

func main() {
	fmt.Println(findSpecialInteger([]int{1, 1, 2, 2, 3, 3, 3, 3}))
	fmt.Println(findSpecialInteger([]int{1, 2, 2, 6, 6, 6, 6, 7, 10}))
	fmt.Println(findSpecialInteger([]int{1, 2, 2, 6, 7}))
	fmt.Println(findSpecialInteger([]int{1, 2, 2, 6}))
	fmt.Println(findSpecialInteger([]int{1, 2, 2}))
	fmt.Println(findSpecialInteger([]int{1, 2}))
	fmt.Println(findSpecialInteger([]int{1}))
}
