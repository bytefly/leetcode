package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	p, q := l1, l2
	dummy := &ListNode{}
	s := dummy
	for p != nil && q != nil {
		if p.Val >= q.Val {
			s.Next = &ListNode{q.Val, nil}
			q = q.Next
		} else {
			s.Next = &ListNode{p.Val, nil}
			p = p.Next
		}
		s = s.Next
	}
	for p != nil {
		s.Next = &ListNode{p.Val, nil}
		s = s.Next
		p = p.Next
	}
	for q != nil {
		s.Next = &ListNode{q.Val, nil}
		s = s.Next
		q = q.Next
	}
	return dummy.Next
}

func main() {
	l1 := &ListNode{1, nil}
	node1 := &ListNode{2, nil}
	node2 := &ListNode{4, nil}
	l1.Next = node1
	node1.Next = node2

	l2 := &ListNode{1, nil}
	node1 = &ListNode{3, nil}
	node2 = &ListNode{4, nil}
	l2.Next = node1
	node1.Next = node2

	l3 := mergeTwoLists(l1, l2)
	p := l3
	for p != nil {
		fmt.Print(p.Val, " ")
		p = p.Next
	}
	fmt.Println()
}
