package main

import (
	"fmt"
	"strconv"
)

func confusingNumber(N int) bool {
	m := [10]bool{true, true, false, false, false, false, true, false, true, true}
	num := strconv.Itoa(N)
	left, right, cnt := 0, len(num)-1, 0

	if N == 6 || N == 9 {
		return true
	}
	for left < right {
		//invalid digit
		if !m[num[left]-'0'] || !m[num[right]-'0'] {
			return false
		}

		if num[left] == num[right] || (num[left] == '6' && num[right] == '9') || (num[left] == '9' && num[right] == '6') {
			cnt++
		}
		left++
		right--
	}

	//invalid middle digit
	if left == right && !m[num[left]-'0'] {
		return false
	}
	//same after rotate
	if cnt == len(num)/2 {
		return false
	}
	return true
}

func main() {
	fmt.Println(confusingNumber(6))
	fmt.Println(confusingNumber(1))
	fmt.Println(confusingNumber(2))
	fmt.Println(confusingNumber(150))
	fmt.Println(confusingNumber(6169))
	fmt.Println(confusingNumber(619))
	fmt.Println(confusingNumber(69))
	fmt.Println(confusingNumber(88))
	fmt.Println(confusingNumber(01))
	fmt.Println(confusingNumber(11))
	fmt.Println(confusingNumber(808))
	fmt.Println(confusingNumber(962))
}
