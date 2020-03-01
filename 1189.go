package main

import "fmt"

func maxNumberOfBalloons(text string) int {
	var aCnt, bCnt, lCnt, oCnt, nCnt int
	for _, c := range text {
		switch c {
		case 'a':
			aCnt++
		case 'b':
			bCnt++
		case 'l':
			lCnt++
		case 'o':
			oCnt++
		case 'n':
			nCnt++
		}
	}

	min := bCnt
	if aCnt < min {
		min = aCnt
	}
	if lCnt/2 < min {
		min = lCnt / 2
	}
	if oCnt/2 < min {
		min = oCnt / 2
	}
	if nCnt < min {
		min = nCnt
	}
	return min
}

func main() {
	fmt.Println(maxNumberOfBalloons("nlaebolko"))
	fmt.Println(maxNumberOfBalloons("loonbalxballpoon"))
	fmt.Println(maxNumberOfBalloons("leetcode"))
}
