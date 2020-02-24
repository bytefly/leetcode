package main

import "fmt"

func replaceElements(arr []int) []int {
	ans := make([]int, len(arr))
	ans[len(ans)-1] = -1
	for i := len(ans) - 2; i >= 0; i-- {
		ans[i] = arr[i+1]
		if ans[i+1] > ans[i] {
			ans[i] = ans[i+1]
		}
	}
	return ans
}

func main() {
	fmt.Println(replaceElements([]int{17}))
	fmt.Println(replaceElements([]int{17, 18}))
	fmt.Println(replaceElements([]int{17, 18, 5, 4, 6, 1}))
}
