package main

import (
	"fmt"
	"sort"
)

type KthLargest struct {
	K    int
	Nums []int
}

func Constructor(k int, nums []int) KthLargest {
	result := KthLargest{k, make([]int, len(nums))}
	copy(result.Nums, nums)
	sort.Ints(result.Nums)
	return result
}

func (this *KthLargest) Add(val int) int {
	left, right := 0, len(this.Nums)-1
	var pos int
	for left <= right {
		mid := left + (right-left)/2
		if this.Nums[mid] == val {
			pos = mid
			break
		}
		if this.Nums[mid] > val {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if left > right {
		pos = left
	}

	this.Nums = append(this.Nums[:pos], append([]int{val}, this.Nums[pos:]...)...)
	if this.K > len(this.Nums) {
		return this.Nums[0]
	}
	return this.Nums[len(this.Nums)-this.K]
}

func main() {
	obj := Constructor(1, []int{})
	fmt.Println(obj.Add(3))
	fmt.Println(obj.Add(2))
	fmt.Println(obj.Add(4))

	obj = Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(obj.Add(3))
	fmt.Println(obj.Add(5))
	fmt.Println(obj.Add(10))
	fmt.Println(obj.Add(9))
	fmt.Println(obj.Add(4))
	fmt.Println(obj.Add(1))
}
