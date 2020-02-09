package main

import (
	"fmt"
)

func readBinaryWatch(num int) []string {
	var ans []string
	var j uint
	var hour, minute int //4 bits(&0x0F) and 6 bits(&0x3F)

	for i := 0; i < 1024; i++ {
		no1Num := 0
		for j = 0; j < 10; j++ {
			if (i>>j)&1 == 1 {
				no1Num++
			}
		}
		if no1Num == num {
			hour, minute = i&0xF, (i&0x3F0)>>4
			if hour < 12 && minute < 60 {
				ans = append(ans, fmt.Sprintf("%d:%02d", hour, minute))
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(readBinaryWatch(2))
}
