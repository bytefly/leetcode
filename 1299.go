package main

import "fmt"

func replaceElements(arr []int) []int {
	ans := make([]int, len(arr))
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > ans[i] {
				ans[i] = arr[j]
			}
		}
	}
	ans[len(ans)-1] = -1
	return ans
}

func main() {
	fmt.Println(replaceElements([]int{17}))
	fmt.Println(replaceElements([]int{17, 18}))
	fmt.Println(replaceElements([]int{17, 18, 5, 4, 6, 1}))
}
