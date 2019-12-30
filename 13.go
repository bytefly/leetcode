package main

import (
	"fmt"
)

func romanToInt(s string) int {
	specNum := map[int]string{4: "IV", 9: "IX", 40: "XL", 90: "XC", 400: "CD", 900: "CM"}
	constChar := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	constNum := map[int]byte{1: 'I', 5: 'V', 10: 'X', 50: 'L', 100: 'C', 500: 'D', 1000: 'M'}

	var i, num int
	for i < len(s) {
		var v1, v2 int
		v1, ok1 := constChar[s[i]]
		if !ok1 {
			return -1
		}

		if i < len(s)-1 {
			_, ok2 := constChar[s[i+1]]
			if !ok2 {
				return -1
			}
			v2 = constChar[s[i+1]]
		}

		if v1 >= v2 {
			num += v1
			if i > 0 {
				if _, ok1 := specNum[num]; ok1 { //for IIII
					return -1
				}
				if _, ok1 := constNum[num]; ok1 { //for IVI
					return -1
				}
			}
			i++
			continue
		}

		//handle special number
		t := v2 - v1
		if _, ok1 := specNum[t]; !ok1 {
			return -1
		}
		if i > 0 && num <= t { //for IVIX/IVIV
			return -1
		}

		num += t
		if i > 0 {
			if _, ok1 := specNum[num]; ok1 { //for VIV
				return -1
			}
			if _, ok1 := constNum[num]; ok1 { //for VIIV
				return -1
			}
		}
		i += 2
	}

	return num
}

func main() {
	testcase := map[string]int{
		"IBC":      -1,
		"MID":      -1,
		"IVIV":     -1,
		"XVXV":     -1,
		"VXV":      -1,
		"IIV":      -1,
		"IIII":     -1,
		"I":        1,
		"III":      3,
		"IV":       4,
		"IX":       9,
		"LVIII":    58,
		"MCDLXXVI": 1476,
		"MCMXCIV":  1994,
	}
	for k, v := range testcase {
		m := romanToInt(k)
		if v != m {
			fmt.Println("test for '", k, "' fail, should be:", v, "but get", m)
		}
	}
}
