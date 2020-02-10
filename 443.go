package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	var prevChar byte
	var prevCnt, ans int

	for _, v := range chars {
		if prevChar != 0 && prevChar != v {
			chars[ans] = prevChar
			ans++
			if prevCnt > 1 {
				numStr := strconv.FormatInt(int64(prevCnt), 10)
				for j := 0; j < len(numStr); j++ {
					chars[ans] = numStr[j]
					ans++
				}
			}

			prevChar = v
			prevCnt = 1
		} else {
			//handle the fist byte
			if prevChar == 0 {
				prevChar = v
			}
			prevCnt++
		}
	}
	//handle same chars in the tail
	chars[ans] = prevChar
	ans++
	if prevCnt > 1 {
		numStr := strconv.FormatInt(int64(prevCnt), 10)
		for j := 0; j < len(numStr); j++ {
			chars[ans] = numStr[j]
			ans++
		}
	}

	if ans < len(chars) {
		chars = chars[:ans]
	}
	return ans
}

func main() {
	s := []byte("a")
	fmt.Println(compress(s), string(s))
	s = []byte("aab")
	fmt.Println(compress(s), string(s))
	s = []byte("aabb")
	fmt.Println(compress(s), string(s))
	s = []byte("abbbbbbbbbbbb")
	fmt.Println(compress(s), string(s))
	s = []byte("aabbccc")
	fmt.Println(compress(s), string(s))
	s = []byte("aabbcccb")
	fmt.Println(compress(s), string(s))
}
