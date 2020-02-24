package main

import "fmt"

func replaceElements(arr []int) []int {
	ans := make([]int, len(arr))
	for i := 0; i < len(arr)-1; i++ {
		max := 0
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > max {
				max = arr[j]
			}
		}
		ans[i] = max
	}
	ans[len(ans)-1] = -1
	return ans
}

func main() {
	fmt.Println(replaceElements([]int{17}))
	fmt.Println(replaceElements([]int{17, 18}))
	fmt.Println(replaceElements([]int{17, 18, 5, 4, 6, 1}))
}
