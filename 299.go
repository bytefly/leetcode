package main

import (
	"fmt"
)

func getHint(secret string, guess string) string {
	state := make([]bool, len(secret))
	var bullCnt, cowCnt int
	for i := 0; i < len(guess); i++ {
		if secret[i] == guess[i] {
			bullCnt++
			state[i] = true
		}
	}
	for i := 0; i < len(guess); i++ {
		for j := 0; j < len(secret); j++ {
			//mark the first position
			if secret[i] != guess[i] && secret[j] == guess[i] && !state[j] {
				cowCnt++
				state[j] = true
				break
			}
		}
	}

	return fmt.Sprintf("%dA%dB", bullCnt, cowCnt)
}

func main() {
	fmt.Println(getHint("1122", "1222"))
	fmt.Println(getHint("1807", "7810"))
	fmt.Println(getHint("1123", "0111"))
	fmt.Println(getHint("112311", "011134"))
}
