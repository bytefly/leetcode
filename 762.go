package main

import (
	"fmt"
)

func countPrimeSetBits(L int, R int) int {
	var ans int
	for L <= R {
		n, cnt := L, 0
		for n > 0 {
			if n&1 == 1 {
				cnt++
			}
			n >>= 1
		}

		if cnt > 1 {
			ans++
			for n = 2; n*n <= cnt; n++ {
				if cnt%n == 0 {
					ans--
					break
				}
			}
		}
		L++
	}

	return ans
}

func main() {
	fmt.Println(countPrimeSetBits(6, 10))
	fmt.Println(countPrimeSetBits(10, 15))
}
