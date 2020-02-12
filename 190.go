package main

import (
	"fmt"
	"strconv"
)

func reverseBits(num uint32) uint32 {
	str := fmt.Sprintf("%032b", num)
	ans := make([]byte, len(str))
	for i := 0; i < len(str); i++ {
		ans[i] = str[len(str)-1-i]
	}
	t, _ := strconv.ParseUint(string(ans), 2, 32)
	return uint32(t)
}

func main() {
	fmt.Println(reverseBits(0))
	fmt.Println(reverseBits(1))
	fmt.Println(reverseBits(43261596))
	fmt.Println(reverseBits(4294967293))
}
