package main

import "fmt"

func duplicateZeros(arr []int) {
	var start, length int
	var ignoreLastZero bool
	for start = 0; start < len(arr); start++ {
		length++
		if arr[start] == 0 {
			if length == len(arr) {
				ignoreLastZero = true
				break
			}
			length++
		}
		if length == len(arr) {
			break
		}
	}
	if start == len(arr) {
		start--
	}

	j := len(arr) - 1
	for i := start; i >= 0; i-- {
		arr[j] = arr[i]
		j--

		if arr[i] == 0 {
			if ignoreLastZero {
				ignoreLastZero = false
				continue
			}
			arr[j] = arr[i]
			j--
		}
	}
}

func main() {
	arr := []int{1, 0, 2, 3, 0, 4, 5, 0}
	duplicateZeros(arr)
	fmt.Println(arr)

	arr = []int{1, 0, 2, 3, 0, 0, 5, 0}
	duplicateZeros(arr)
	fmt.Println(arr)

	arr = []int{1, 2, 3}
	duplicateZeros(arr)
	fmt.Println(arr)

	arr = []int{0, 0, 0, 0, 0, 0, 0, 0}
	duplicateZeros(arr)
	fmt.Println(arr)
}
