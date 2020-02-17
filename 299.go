package main

import (
	"fmt"
)

func getHint(secret string, guess string) string {
	state := make(map[byte]int)
	var bullCnt, cowCnt int
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bullCnt++
		}
		state[secret[i]]++
	}
	for i := 0; i < len(guess); i++ {
		if v, ok := state[guess[i]]; ok && v > 0 {
			cowCnt++
			state[guess[i]]--
		}
	}
	cowCnt -= bullCnt

	return fmt.Sprintf("%dA%dB", bullCnt, cowCnt)
}

func main() {
	fmt.Println(getHint("1122", "1222"))
	fmt.Println(getHint("1807", "7810"))
	fmt.Println(getHint("1123", "0111"))
	fmt.Println(getHint("112311", "011134"))
}
