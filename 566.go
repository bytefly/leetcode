package main

import (
	"fmt"
)

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums) == 0 || len(nums)*len(nums[0]) != r*c {
		return nums
	}

	var newNums []int
	ans := make([][]int, r)
	for _, h := range nums {
		for _, v := range h {
			newNums = append(newNums, v)
		}
	}

	for i := 0; i < r; i++ {
		ans[i] = newNums[i*c : i*c+c]
	}
	return ans
}

func main() {
	fmt.Println(matrixReshape([][]int{[]int{1, 2}, []int{3, 4}}, 1, 4))
	fmt.Println(matrixReshape([][]int{[]int{1, 2}, []int{3, 4}}, 4, 2))
	fmt.Println(matrixReshape([][]int{[]int{1, 2}, []int{3, 4}}, 4, 1))
	fmt.Println(matrixReshape([][]int{[]int{1, 2, 3}, []int{4, 5, 6}}, 3, 2))
}
