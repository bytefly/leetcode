package main

import (
	"fmt"
)

type NumArray struct {
	sum []int
}

func Constructor(nums []int) NumArray {
	if len(nums) == 0 {
		return NumArray{[]int{}}
	}

	var sum []int
	sum = append(sum, nums[0])
	for i := 1; i < len(nums); i++ {
		sum = append(sum, sum[len(sum)-1]+nums[i])
	}
	return NumArray{sum}
}

func (this *NumArray) SumRange(i int, j int) int {
	if i < 0 || i >= len(this.sum) || j < 0 || j >= len(this.sum) {
		return -1
	}
	if i > j {
		i, j = j, i
	}
	if i == 0 {
		return this.sum[j]
	}
	return this.sum[j] - this.sum[i-1]
}

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := Constructor(nums)
	fmt.Println(obj.SumRange(0, 2))
	fmt.Println(obj.SumRange(2, 5))
	fmt.Println(obj.SumRange(0, 5))
}
