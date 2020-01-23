package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	var (
		newHead *ListNode
		length  int
	)
	if head == nil {
		return head
	}

	//find the tail
	p := head
	for p.Next != nil {
		p = p.Next
		length++
	}
	length++
	//make the ring circle
	p.Next = head

	cnt := k % length
	p = head
	i := 0
	for {
		if i == cnt {
			newHead = p
		} else if i > cnt && p.Next == newHead {
			//break the ring circle
			p.Next = nil
			break
		}

		i++
		p = p.Next
	}
	return newHead
}

func main() {
	node0 := &ListNode{0, nil}
	node1 := &ListNode{1, nil}
	node2 := &ListNode{2, nil}
	node0.Next = node1
	node1.Next = node2
	p := rotateRight(node0, 2)

	for p != nil {
		fmt.Print(p.Val, " ")
		p = p.Next
	}
	fmt.Println()
}
