package main

import (
	"fmt"
	"strconv"
)

func toHexspeak(num string) string {
	n, _ := strconv.Atoi(num)
	b := []byte(fmt.Sprintf("%X", n))
	for i := 0; i < len(b); i++ {
		switch {
		case b[i] == '0':
			b[i] = 'O'
		case b[i] == '1':
			b[i] = 'I'
		case b[i] >= '2' && b[i] <= '9':
			return "ERROR"
		}
	}
	return string(b)
}

func main() {
	fmt.Println(toHexspeak("257"))
	fmt.Println(toHexspeak("3"))
}
