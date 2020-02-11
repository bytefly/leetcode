package main

import (
	"fmt"
)

func checkRecord(s string) bool {
	var aNum, lNum int

	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			aNum++
		} else if s[i] == 'L' {
			lNum++
		}

		if lNum > 2 || aNum > 1 {
			return false
		}
		if s[i] != 'L' && lNum != 0 {
			lNum = 0
		}
	}

	return true
}

func main() {
	fmt.Println(checkRecord("PPALLP"))
	fmt.Println(checkRecord("PPALLL"))
	fmt.Println(checkRecord("PPALLA"))
	fmt.Println(checkRecord("LPPLLA"))
}
