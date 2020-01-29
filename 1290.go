package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	var num int
	p := head
	for p != nil {
		num = num * 2
		num += p.Val
		p = p.Next
	}

	return num
}

func createListNode(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	var head, p *ListNode
	for i := 0; i < len(vals); i++ {
		node := &ListNode{vals[i], nil}
		if i == 0 {
			head = node
		} else {
			p.Next = node
		}
		p = node
	}
	return head
}

func main() {
	fmt.Println(getDecimalValue(createListNode([]int{0})))
	fmt.Println(getDecimalValue(createListNode([]int{1})))
	fmt.Println(getDecimalValue(createListNode([]int{1, 1})))
	fmt.Println(getDecimalValue(createListNode([]int{1, 1, 1})))
	fmt.Println(getDecimalValue(createListNode([]int{1, 1, 1, 1})))
	fmt.Println(getDecimalValue(createListNode([]int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0})))
}
