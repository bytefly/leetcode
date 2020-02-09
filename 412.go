package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	var ans []string

	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				ans = append(ans, "FizzBuzz")
			} else {
				ans = append(ans, "Fizz")
			}
		} else if i%5 == 0 {
			ans = append(ans, "Buzz")
		} else {
			ans = append(ans, strconv.FormatInt(int64(i), 10))
		}
	}
	return ans
}

func main() {
	fmt.Println(fizzBuzz(15))
	fmt.Println(fizzBuzz(-15))
}
