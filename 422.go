package main

import "fmt"

func validWordSquare(words []string) bool {
	for i := 0; i < len(words); i++ { // row index
		for j := 0; j < len(words[i]); j++ { //column index
			//a string is too long, rows are short
			if j == len(words) {
				return false
			}
			//a string is too long, columns are short
			if i >= len(words[j]) {
				return false
			}
			if words[j][i] != words[i][j] {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Println(validWordSquare([]string{"abcd", "bnrt", "crmy", "dtye"}))
	fmt.Println(validWordSquare([]string{"abcd", "bnrt", "crm", "dt"}))
	fmt.Println(validWordSquare([]string{"ball", "area", "read", "lady"}))
	fmt.Println(validWordSquare([]string{"ball", "asee", "let", "lep"}))
	fmt.Println(validWordSquare([]string{"ball", "asee", "letp", "le"}))
	fmt.Println(validWordSquare([]string{"abc", "b"}))
	fmt.Println(validWordSquare([]string{"ab", "b", "c"}))
	fmt.Println(validWordSquare([]string{"abcd", "bef", "cfga", "d"}))
}
