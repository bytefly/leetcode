package main

import "fmt"

func findSpecialInteger(arr []int) int {
	ans := arr[0]
	cnt, maxCnt := 1, 1

	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			cnt++
			if cnt > maxCnt {
				ans = arr[i]
				maxCnt = cnt
			}
		} else {
			cnt = 1
		}
	}

	return ans
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
