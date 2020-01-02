package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	p, q := dummy, dummy

	for i := 0; i < n+1; i++ {
		if q == nil {
			break
		}
		q = q.Next
	}

	for q != nil {
		p = p.Next
		q = q.Next
	}

	p.Next = p.Next.Next
	head = dummy.Next

	return head
}

func main() {
	node0 := &ListNode{1, nil}
	node1 := &ListNode{2, nil}
	node2 := &ListNode{3, nil}
	node3 := &ListNode{4, nil}
	node4 := &ListNode{5, nil}
	node0.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	node := removeNthFromEnd(node0, 1)
	head := node
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}
