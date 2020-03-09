package main

import "fmt"

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{[]int{nums[0]}}
	}

	var ans [][]int
	for i := 0; i < len(nums); i++ {
		t := permute(append(append([]int{}, nums[:i]...), nums[i+1:]...))
		for _, s := range t {
			s = append(s, nums[i])
			ans = append(ans, s)
		}
	}
	return ans
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
