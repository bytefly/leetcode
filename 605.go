package main

import (
	"fmt"
)

func canPlaceFlowers(flowerbed []int, n int) bool {
	var maxN int
	start := -1
	for i, v := range flowerbed {
		if v == 1 {
			//beginning with 0s
			if start < 0 {
				maxN += (i - start - 1) / 2
			} else {
				maxN += (i - start - 2) / 2
			}
			start = i
		}
	}
	//ending with 0s
	if start < len(flowerbed)-1 {
		//all 0s
		if start < 0 {
			maxN += (len(flowerbed) + 1) / 2
		} else {
			maxN += (len(flowerbed) - 1 - start) / 2
		}
	}

	return maxN >= n
}

func main() {
	fmt.Println(canPlaceFlowers([]int{0}, 1))
	fmt.Println(canPlaceFlowers([]int{0, 0, 0, 0, 0}, 1))

	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 0, 1}, 2))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 0, 0, 1}, 2))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 0, 0, 0, 1}, 3))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 0, 0, 0, 0, 1}, 3))

	fmt.Println(canPlaceFlowers([]int{1, 0, 1, 0}, 1))
	fmt.Println(canPlaceFlowers([]int{1, 0, 1, 0, 0}, 2))
	fmt.Println(canPlaceFlowers([]int{1, 0, 1, 0, 0, 0}, 2))

	fmt.Println(canPlaceFlowers([]int{0, 1, 0, 1}, 2))
	fmt.Println(canPlaceFlowers([]int{0, 0, 1, 0, 1}, 2))
	fmt.Println(canPlaceFlowers([]int{0, 0, 0, 1, 0, 1}, 2))
	fmt.Println(canPlaceFlowers([]int{0, 0, 0, 0, 1, 0, 1}, 2))
}
