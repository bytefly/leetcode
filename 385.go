package main

import (
	"fmt"
	"strconv"
	"strings"
)

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
	Val  int
	List []*NestedInteger
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool {
	if n.List == nil {
		return true
	}
	return false
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() int {
	return n.Val
}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {
	n.Val = value
}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) {
	if n.List == nil {
		n.List = make([]*NestedInteger, 0)
	}
	n.List = append(n.List, &elem)
}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []*NestedInteger {
	return n.List
}

func (n NestedInteger) String() string {
	if n.IsInteger() {
		return strconv.FormatInt(int64(n.GetInteger()), 10)
	}

	list := n.GetList()
	str := "["
	for i := 0; i < len(list); i++ {
		str += list[i].String()
		//str += "("+fmt.Sprintf("%p", list[i])+")"
		if i != len(list)-1 {
			str += ","
		}
	}
	str += "]"
	return str
}

func deserialize(s string) *NestedInteger {
	var ans *NestedInteger
	var val int

	if strings.IndexByte(s, '[') == -1 {
		val, _ = strconv.Atoi(s)
		ans = new(NestedInteger)
		ans.SetInteger(val)
		return ans
	}

	stack := make([]*NestedInteger, 0)
	var num string
	for i := 0; i < len(s); i++ {
		switch {
		case s[i] == '[':
			var q NestedInteger
			if len(stack) == 0 {
				ans = &q
				stack = append(stack, &q)
			} else {
				p := stack[len(stack)-1]
				p.Add(q)
				list := p.GetList()
				stack = append(stack, list[len(list)-1])
			}

		case s[i] == ']':
			if num != "" {
				p := stack[len(stack)-1]
				var q NestedInteger
				val, _ = strconv.Atoi(num)
				q.SetInteger(val)
				p.Add(q)
				num = ""
			}
			stack = stack[:len(stack)-1]

		case s[i] == '-' || (s[i] >= '0' && s[i] <= '9'):
			if num == "" {
				num = string(s[i])
			} else {
				num += string(s[i])
			}

		case s[i] == ',':
			if num != "" {
				p := stack[len(stack)-1]
				var q NestedInteger
				val, _ = strconv.Atoi(num)
				q.SetInteger(val)
				p.Add(q)
				num = ""
			}
		}
	}

	return ans
}

func main() {
	ans := deserialize("[123]")
	fmt.Println(ans)
	ans = deserialize("324")
	fmt.Println(ans)
	ans = deserialize("[123,[234]]")
	fmt.Println(ans)
	ans = deserialize("[123,[456,[789]]]")
	fmt.Println(ans)
	ans = deserialize("[[123],[456,[789]]]")
	fmt.Println(ans)
}
