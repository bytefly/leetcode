package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var p *ListNode
	q := head
	for q != nil {
		t := q.Next
		q.Next = p
		p = q
		q = t
	}
	return p
}

func main() {
	node0 := &ListNode{1, nil}
	node1 := &ListNode{2, nil}
	node2 := &ListNode{3, nil}
	node0.Next = node1
	node1.Next = node2
	p := reverseList(node0)
	q := p
	for q != nil {
		fmt.Print(q.Val, " ")
		q = q.Next
	}
	fmt.Println()
}
