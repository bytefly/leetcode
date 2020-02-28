package main

import "fmt"

func transformArray(arr []int) []int {
	for {
		cnt := 0
		prev := arr[0]
		for i := 1; i < len(arr)-1; i++ {
			t := arr[i]
			if arr[i] < prev && arr[i] < arr[i+1] {
				arr[i]++
				cnt++
			} else if arr[i] > prev && arr[i] > arr[i+1] {
				arr[i]--
				cnt++
			}
			prev = t
		}

		if cnt == 0 {
			break
		}
	}
	return arr
}

func main() {
	fmt.Println(transformArray([]int{6, 2, 3, 4}))
	fmt.Println(transformArray([]int{1, 6, 3, 4, 3, 5}))
	fmt.Println(transformArray([]int{2, 1, 2, 1, 1, 2, 2, 1}))
}
