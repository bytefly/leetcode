package main

import "fmt"

func numPairsDivisibleBy60(time []int) int {
	var ans int
	var m [60]int
	for i := 0; i < len(time); i++ {
		m[time[i]%60]++
	}

	for k, _ := range m {
		if k == 30 || k == 0 {
			if m[k] > 0 {
				ans += m[k] * (m[k] - 1) / 2
			}
		} else {
			if k < 30 && m[60-k] > 0 {
				ans += m[k] * m[60-k]
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(numPairsDivisibleBy60([]int{30, 20, 150, 100, 40}))
	fmt.Println(numPairsDivisibleBy60([]int{60, 60, 60}))
}
