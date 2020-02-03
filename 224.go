package main

import (
	"fmt"
	"strconv"
)

func calculate(s string) int {
	var number1, number2 int
	var top, operator string
	stack := make([]string, 0)

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ' ':
			//ignored
		case '(':
			stack = append(stack, "(")
		case ')':
			//popup "(" and push the number
			top = stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			//calcutlate common expression
			if len(stack) > 0 && (stack[len(stack)-1] == "+" || stack[len(stack)-1] == "-") {
				number1, _ = strconv.Atoi(stack[len(stack)-2])
				number2, _ = strconv.Atoi(top)
				operator = stack[len(stack)-1]
				stack = stack[:len(stack)-2]

				//push the result
				if operator == "+" {
					stack = append(stack, strconv.FormatInt(int64(number1+number2), 10))
				} else {
					stack = append(stack, strconv.FormatInt(int64(number1-number2), 10))
				}
			} else {
				stack = append(stack, top)
			}
		case '+':
			fallthrough
		case '-':
			//push the operator
			stack = append(stack, string(s[i]))
		default: //digit number
			if i == len(s)-1 || (s[i+1] < '0' || s[i+1] > '9') {
				//calcutlate common expression
				if len(stack) >= 3 && stack[len(stack)-1] != "+" && stack[len(stack)-1] != "-" && stack[len(stack)-1] != "(" &&
					(stack[len(stack)-2] == "+" || stack[len(stack)-2] == "-") {
					number1, _ = strconv.Atoi(stack[len(stack)-3])
					number2, _ = strconv.Atoi(stack[len(stack)-1] + string(s[i]))
					operator = stack[len(stack)-2]
					stack = stack[:len(stack)-3]

					//push the result
					if operator == "+" {
						stack = append(stack, strconv.FormatInt(int64(number1+number2), 10))
					} else {
						stack = append(stack, strconv.FormatInt(int64(number1-number2), 10))
					}
					break
				} else if len(stack) >= 2 && stack[len(stack)-2] != "+" && stack[len(stack)-2] != "-" && stack[len(stack)-2] != "(" &&
					(stack[len(stack)-1] == "+" || stack[len(stack)-1] == "-") {
					number1, _ = strconv.Atoi(stack[len(stack)-2])
					number2, _ = strconv.Atoi(string(s[i]))
					operator = stack[len(stack)-1]
					stack = stack[:len(stack)-2]

					//push the result
					if operator == "+" {
						stack = append(stack, strconv.FormatInt(int64(number1+number2), 10))
					} else {
						stack = append(stack, strconv.FormatInt(int64(number1-number2), 10))
					}
					break
				}
			}
			//push the digit
			if len(stack) > 0 && stack[len(stack)-1] != "+" &&
				stack[len(stack)-1] != "-" && stack[len(stack)-1] != "(" {
				top = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, top+string(s[i]))
			} else {
				stack = append(stack, string(s[i]))
			}
		}
	}

	ans, _ := strconv.ParseInt(stack[0], 10, 32)
	return int(ans)
}

func main() {
	fmt.Println(calculate("1 + 1"))
	fmt.Println(calculate("((1 + 1))"))
	fmt.Println(calculate(" 2-1 + 2 "))
	fmt.Println(calculate(" 4-5 + 2 "))
	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
	fmt.Println(calculate("13 + 1 + 345"))
	fmt.Println(calculate("(10+(40+50+20)-30)+(60+80)"))
}
