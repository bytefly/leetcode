package main

import (
	"fmt"
	"math"
)

func arrangeCoins(n int) int {
	t := int(math.Sqrt(float64(n << 1)))
	if t*t+t > n<<1 {
		t--
	}
	return t
}

func main() {
	fmt.Println(arrangeCoins(0))
	fmt.Println(arrangeCoins(1))
	fmt.Println(arrangeCoins(2))
	fmt.Println(arrangeCoins(3))
	fmt.Println(arrangeCoins(5))
	fmt.Println(arrangeCoins(6))
	fmt.Println(arrangeCoins(8))
	fmt.Println(arrangeCoins(9))
	fmt.Println(arrangeCoins(10))
}
