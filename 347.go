package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	var ans []int
	var prev, cnt int

	n := len(nums)
	m := make(map[int][]int)

	for i := n/2 - 1; i >= 0; i-- {
		heaplize(nums, i, n)
	}

	for i := n - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		//check and fill the map
		if cnt == 0 {
			prev = nums[i]
			cnt = 1
		} else if nums[i] != prev {
			m[cnt] = append(m[cnt], prev)
			prev = nums[i]
			cnt = 1
		} else {
			cnt++
		}
		heaplize(nums, 0, i)
	}
	m[cnt] = append(m[cnt], prev)

	//get the result from the most count key
	for i := n; i > 0; i-- {
		if v, ok := m[i]; ok {
			needCopy := k - len(ans)
			if needCopy >= len(v) {
				ans = append(ans, v...)
			} else {
				ans = append(ans, v[:needCopy]...)
			}
		}
	}

	return ans
}

func heaplize(nums []int, i, end int) {
	left, right := 2*i+1, 2*i+2
	large := i
	if left < end && nums[left] > nums[i] {
		large = left
	}
	if right < end && nums[right] > nums[large] {
		large = right
	}

	if large != i {
		nums[i], nums[large] = nums[large], nums[i]
		heaplize(nums, large, end)
	}
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent([]int{1}, 1))
}
