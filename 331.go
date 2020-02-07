package main

import (
	"fmt"
	"strings"
)

func isValidSerialization(preorder string) bool {
	var i, stopIdx int
	vals := strings.Split(preorder, ",")
	stack := make([]string, 0)

	if preorder == "" {
		return true
	}
	for i = 0; i < len(vals); i++ {
		//string prefix with # is allowed and ignored
		if i == 0 && vals[0] == "#" {
			continue
		}
		stack = append(stack, vals[i])
		if vals[i] == "#" {
			stopIdx = -1
			//loop check valid node(both left and right nodes are '#')
			//then mark it as '#' and remove child nodes
			for j := len(stack) - 1; j > 0; {
				if j >= 2 && stack[j] == "#" && stack[j-1] == "#" && stack[j-2] != "#" {
					stack[j-2] = "#"
					j -= 2
					stopIdx = j
				} else {
					stopIdx = j + 1
					break
				}
			}
			if stopIdx >= 0 {
				stack = stack[:stopIdx]
			}
		}
		//move i forward and jump out if stack is empty
		if len(stack) == 0 {
			i++
			break
		}
	}

	//only valid when all nodes handled
	return len(stack) == 0 && i == len(vals)
}

func main() {
	cases := map[string]bool{
		"9,3,4,#,#,1,#,#,#,2,#,6,#,#": false,
		"9,3,4,#,#,1,#,#,2,#,6,#,#":   true,
		"1,#,#":                       true,
		"1,#":                         false,
		"9,#,#,1":                     false,
		"":                            true,
		"#":                           true,
		"#,#":                         false,
		"#,1":                         false,
	}
	for k, v := range cases {
		if isValidSerialization(k) != v {
			fmt.Println("test failed for", k)
		}
	}
}
