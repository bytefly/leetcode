package main

import "fmt"

func bitwiseComplement(N int) int {
	var pos uint
	if N == 0 {
		return 1
	}
	for pos = 0; pos < 32; pos++ {
		if N <= (1 << pos) {
			if N != (1 << pos) {
				pos -= 1
			}
			break
		}
	}

	return (1<<(pos+1) - 1) & ^N
}

func main() {
	fmt.Println(bitwiseComplement(5))
	fmt.Println(bitwiseComplement(7))
	fmt.Println(bitwiseComplement(10))
}
