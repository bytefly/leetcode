package main

import "fmt"

func divisorGame(N int) bool {
	return N%2 == 0
}

func main() {
	fmt.Println(divisorGame(2))
	fmt.Println(divisorGame(3))
}
