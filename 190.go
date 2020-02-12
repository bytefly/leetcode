package main

import (
	"fmt"
)

func reverseBits(num uint32) uint32 {
	var ans uint32
	for i := 0; i < 32; i++ {
		ans *= 2
		ans += num % 2

		num /= 2
	}

	return ans
}

func main() {
	fmt.Println(reverseBits(0))
	fmt.Println(reverseBits(1))
	fmt.Println(reverseBits(43261596))
	fmt.Println(reverseBits(4294967293))
}
