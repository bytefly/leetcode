package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	stack := make([]string, 0)
	tokens := strings.Split(path, "/")
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "":
		case ".":
			//do nothing
		case "..":
			//jump to parent folder
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			//push to stack
			stack = append(stack, tokens[i])
		}
	}

	if len(stack) == 0 {
		return "/"
	}

	var newPath string
	for i := 0; i < len(stack); i++ {
		newPath += "/" + stack[i]
	}
	return newPath
}

func main() {
	tests := []string{"/home/", "/home", "/../", "/", "/home//foo/", "/home/foo",
		"/a/./b/../../c/", "/c", "/a/../../b/../c//.//", "/c",
		"/a//b////c/d//././/..", "/a/b/c"}
	for i := 0; i < len(tests); i += 2 {
		v := simplifyPath(tests[i])
		if v != tests[i+1] {
			fmt.Println("test for", tests[i], "fail, get:", v, "should be:", tests[i+1])
		}
	}
}
