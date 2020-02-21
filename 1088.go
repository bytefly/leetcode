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

		if (num[left] == num[right] && num[left] != '6' && num[left] != '9') ||
			(num[left] == '6' && num[right] == '9') || (num[left] == '9' && num[right] == '6') {
			cnt++
		}
		left++
		right--
	}

	if left == right {
		//invalid middle digit
		if !m[num[left]-'0'] {
			return false
		}
		//middle digit is 6 or 9
		if num[left] == '6' || num[left] == '9' {
			return true
		}
	}
	//same after rotate
	if cnt == len(num)/2 {
		return false
	}
	return true
}

func confusingNumberII(N int) int {
	var ans int
	for i := 1; i <= N; i++ {
		if confusingNumber(i) {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(confusingNumberII(20))
	fmt.Println(confusingNumberII(100))
	fmt.Println(confusingNumberII(86))
	fmt.Println(confusingNumberII(195))
}
