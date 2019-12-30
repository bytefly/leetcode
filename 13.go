package main

import (
	"fmt"
)

func romanToInt(s string) int {
	specStr := map[int]string{4: "IV", 9: "IX", 40: "XL", 90: "XC", 400: "CD", 900: "CM"}
	constChar := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	constNum := map[int]byte{1: 'I', 5: 'V', 10: 'X', 50: 'L', 100: 'C', 500: 'D', 1000: 'M'}

	var ok1, ok2 bool
	var num, v1, v2, t int
	for i := len(s) - 1; i >= 0; {
		v1, ok1 = constChar[s[i]]
		if i > 0 {
			v2, ok2 = constChar[s[i-1]]
		}
		//invalid character
		if !ok1 || (i > 0 && !ok2) {
			return -1
		}

		//fmt.Println(i, v1, v2, num)
		if i == 0 {
			if v1 > num || s[i] == s[i+1] {
				if len(s) > 1 { //ignore single char check
					_, ok1 := specStr[num+v1]
					_, ok2 := constNum[num+v1]
					if ok1 || ok2 {
						return -1
					}
				}

				num += v1
			} else {
				num = -1
			}
			break
		}

		if v2 >= v1 {
			t = v1 + v2
			_, ok1 := specStr[t]
			_, ok2 := constNum[t]
			if ok1 || ok2 {
				return -1
			}

			if t > num {
				num += t
			} else {
				return -1
			}
		}
		if v1 > v2 {
			t = v1 - v2
			_, ok1 := specStr[t]
			//not the special num in the map or same with prev special integer
			if ok1 {
				if t > num {
					num += t
				} else {
					return -1
				}
			} else {
				return -1
			}
		}

		i -= 2
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
		"MCDLXXVI": 1676,
		"MCMXCIV":  1994,
	}
	for k, v := range testcase {
		m := romanToInt(k)
		if v != m {
			fmt.Println("test for '", k, "' fail, should be:", v, "but get", m)
		}
	}
}
