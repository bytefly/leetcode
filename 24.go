package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy, p := head.Next, head
	q := p.Next
	var prev *ListNode

	if dummy == nil {
		return head
	}
	for p != nil && q != nil {
		//swap p and q
		t := q.Next
		q.Next = p
		p.Next = t

		//link prev and q
		if prev == nil {
			prev = p
			dummy = q
		} else {
			prev.Next = q
		}

		//move prev,p,q to the next position
		prev = p
		p = prev.Next
		if p != nil {
			q = p.Next
		}
	}
	return dummy
}

func main() {
	node := []int{1, 2, 3, 4, 5}
	var head, q *ListNode
	for i, val := range node {
		n := &ListNode{val, nil}
		if i == 0 {
			head, q = n, n
		} else {
			q.Next = n
			q = q.Next
		}
	}

	head = swapPairs(head)

	q = head
	for q != nil {
		fmt.Print(q.Val, " ")
		q = q.Next
	}
	fmt.Println()
}
