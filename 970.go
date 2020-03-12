package main

import "fmt"

func powerfulIntegers(x int, y int, bound int) []int {
	var ans []int
	m := make(map[int]bool)
	powerX, powerY := 1, 1

	for powerX < bound {
		for powerY < bound {
			t := powerX + powerY
			if t <= bound {
				m[t] = true
			}
			powerY *= y
			if powerY == 1 {
				break
			}
		}

		powerX *= x
		powerY = 1
		if powerX == 1 {
			break
		}
	}

	for k, _ := range m {
		ans = append(ans, k)
	}
	return ans
}

func main() {
	fmt.Println(powerfulIntegers(2, 3, 10))
	fmt.Println(powerfulIntegers(3, 5, 15))
	fmt.Println(powerfulIntegers(2, 1, 10))
	fmt.Println(powerfulIntegers(1, 2, 10))
	fmt.Println(powerfulIntegers(1, 1, 10))
}
