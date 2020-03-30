package main

import (
	"fmt"
	"strconv"
)

func decodeString(s string) string {
	var ans []byte
	var kStack []int
	var strStack []string

	numPos, strPos := -1, -1
	for i := 0; i < len(s); i++ {
		switch {
		case s[i] >= '0' && s[i] <= '9':
			//number followed by string, push the string to stack first, such as: 2[ab3[c]]
			if strPos >= 0 && numPos < 0 {
				strStack = append(strStack, s[strPos:i])
			}
			//mark a start number position
			if numPos < 0 {
				numPos = i
			}
		case (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z'):
			//char after ]
			if strPos < 0 {
				//it's a part of string in stack, such as: ...2[ab2[c]de]...
				if len(strStack) > 0 {
					strStack[len(strStack)-1] += string(s[i])
				} else { //a raw string byte, such as: ...2[ab]cde...
					ans = append(ans, s[i])
				}
			}
		case s[i] == '[':
			//push number to stack and mark string position
			k, _ := strconv.Atoi(s[numPos:i])
			kStack = append(kStack, k)
			numPos = -1
			strPos = i + 1
		case s[i] == ']':
			k := kStack[len(kStack)-1]
			kStack = kStack[:len(kStack)-1]

			var m string
			if strPos >= 0 {
				m = s[strPos:i]
				strPos = -1
			} else {
				m = strStack[len(strStack)-1]
				strStack = strStack[:len(strStack)-1]
			}

			//repeat the string
			var t string
			for i := 0; i < k; i++ {
				t += m
			}
			//add the string to stack or the final result
			if len(strStack) > 0 {
				strStack[len(strStack)-1] += t
			} else {
				ans = append(ans, []byte(t)...)
			}
		}
	}

	return string(ans)
}

func main() {
	fmt.Println(decodeString("3[a10[bc]]"))
	fmt.Println(decodeString("3[a]2[b4[F]c]"))
	fmt.Println(decodeString("3[a]2[bc]"))
	fmt.Println(decodeString("3[a2[c]]"))
	fmt.Println(decodeString("3[a2[c]3[d]]"))
	fmt.Println(decodeString("2[abc]3[cd]ef"))
	fmt.Println(decodeString("2[abc]ef3[c2[d]]"))
}
