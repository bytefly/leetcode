package main

import "fmt"

func maxDistToClosest(seats []int) int {
	var ans, t int

	for i := 0; i < len(seats); i++ {
		if seats[i] == 0 {
			t++
		} else if t > 0 {
			if i-t-1 >= 0 {
				//take care of t = 1, not ok if replaced with t /= 2
				t = t - t/2
			}
			if t > ans {
				ans = t
			}

			t = 0
		}
	}
	//handle elements end with 0s
	if seats[len(seats)-1] == 0 && t > ans {
		ans = t
	}

	return ans
}

func main() {
	fmt.Println(maxDistToClosest([]int{1, 0, 0, 0, 1, 0, 1}))
	fmt.Println(maxDistToClosest([]int{1, 0, 0, 0}))
	fmt.Println(maxDistToClosest([]int{0, 0, 1, 0}))
	fmt.Println(maxDistToClosest([]int{0, 0, 0, 1}))
}
