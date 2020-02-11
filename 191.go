package main

import (
	"fmt"
)

func hammingWeight(num uint32) int {
	var ans int
	for i := 0; i < 32; i++ {
		if num&(1<<uint(i)) != 0 {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(hammingWeight(0))
	fmt.Println(hammingWeight(1))
	fmt.Println(hammingWeight(11))
	fmt.Println(hammingWeight(1024))
	fmt.Println(hammingWeight(4294967293))
}
