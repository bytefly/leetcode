package main

import "fmt"

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
	Val  int
	List []*NestedInteger
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool {
	return len(n.List) == 0
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
	if len(n.List) > 0 {
		n.List = n.List[:0]
	}
}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) {
	n.List = append(n.List, &elem)
}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []*NestedInteger {
	return n.List
}

func depthSum(nestedList []*NestedInteger) int {
	var ans int
	var stack []*NestedInteger

	depth := 1
	for _, ele := range nestedList {
		stack = append(stack, ele)
	}

	for len(stack) > 0 {
		size := len(stack)
		for i := 0; i < size; i++ {
			ele := stack[i]
			if ele.IsInteger() {
				ans += depth * ele.GetInteger()
				continue
			}

			for _, ele2 := range ele.GetList() {
				stack = append(stack, ele2)
			}
		}

		stack = stack[size:]
		depth++
	}

	return ans
}

func main() {
	//[1, [4, [6]]
	i1 := &NestedInteger{1, nil}
	i2 := &NestedInteger{4, nil}
	i3 := &NestedInteger{6, nil}
	i4 := &NestedInteger{0, []*NestedInteger{i3}}
	i5 := &NestedInteger{0, []*NestedInteger{i2, i4}}
	fmt.Println(depthSum([]*NestedInteger{i1, i5}))

	//[[1, 1], 2, [1, 1]]
	i1 = &NestedInteger{1, nil}
	i2 = &NestedInteger{0, []*NestedInteger{i1, i1}}
	i3 = &NestedInteger{2, nil}
	fmt.Println(depthSum([]*NestedInteger{i2, i3, i2}))
}
