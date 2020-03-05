package main

import "fmt"

func duplicateZeros(arr []int) {
	var i, j int
	dup := make([]int, len(arr))
	copy(dup, arr)

	for i < len(arr) {
		arr[i] = dup[j]
		if arr[i] == 0 && i < len(arr)-1 {
			arr[i+1] = 0
			i++
		}

		i++
		j++
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
}
