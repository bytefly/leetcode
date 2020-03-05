package main

import "fmt"

func distributeCandies(candies int, num_people int) []int {
	ans := make([]int, num_people)
	curCandy, j, total := 1, 0, candies

	for {
		if total < curCandy {
			ans[j] += total
			break
		}

		ans[j] += curCandy
		total -= curCandy
		curCandy++
		j++

		if j == num_people {
			j = 0
		}
	}

	return ans
}

func main() {
	fmt.Println(distributeCandies(3, 1))
	fmt.Println(distributeCandies(7, 4))
	fmt.Println(distributeCandies(10, 3))
}
