package main

import (
	"fmt"
	"sort"
)

func sortByBits(arr []int) []int {
	var ans []int
	for _, num := range arr {
		n, cnt := num, 0
		for num > 0 {
			if num&1 > 0 {
				cnt++
			}
			num >>= 1
		}
		//make count as a part of number in ans
		ans = append(ans, (cnt<<16)|n)
	}

	sort.Ints(ans)
	//drop count part
	for i := 0; i < len(ans); i++ {
		ans[i] &= 0xFFFF
	}
	return ans
}

func main() {
	fmt.Println(sortByBits([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}))
	fmt.Println(sortByBits([]int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}))
	fmt.Println(sortByBits([]int{10000, 10000}))
	fmt.Println(sortByBits([]int{2, 3, 5, 7, 11, 13, 17, 19}))
	fmt.Println(sortByBits([]int{10, 100, 1000, 10000}))
}
