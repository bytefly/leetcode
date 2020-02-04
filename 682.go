package main

import (
	"fmt"
	"strconv"
)

func calPoints(ops []string) int {
	var score, total int
	stack := make([]int, 0)
	for i := 0; i < len(ops); i++ {
		switch ops[i] {
		case "C":
			if len(stack) > 0 {
				total -= stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			} else {
				return 0
			}
		case "D":
			if len(stack) > 0 {
				score = stack[len(stack)-1]
				stack = append(stack, score*2)
				total += score * 2
			} else {
				return 0
			}
		case "+":
			if len(stack) > 1 {
				s1, s2 := stack[len(stack)-2], stack[len(stack)-1]
				stack = append(stack, s1+s2)
				total += s1 + s2
			} else {
				return 0
			}
		default:
			score, _ = strconv.Atoi(ops[i])
			stack = append(stack, score)
			total += score
		}
	}

	return total
}

func main() {
	cases := []string{"5", "2", "C", "D", "+"}
	fmt.Println(calPoints(cases))
	cases = []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	fmt.Println(calPoints(cases))
	cases = []string{"5", "+"}
	fmt.Println(calPoints(cases))
}
