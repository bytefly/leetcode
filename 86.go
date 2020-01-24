package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	p := head
	var maxNode, m1 *ListNode
	var minNode, m2 *ListNode

	for p != nil {
		t := &ListNode{p.Val, nil}
		if p.Val >= x {
			if m1 != nil {
				m1.Next = t
			} else {
				maxNode = t
			}
			m1 = t
		} else {
			if m2 != nil {
				m2.Next = t
			} else {
				minNode = t
			}
			m2 = t
		}

		p = p.Next
	}

	if minNode == nil {
		return maxNode
	}

	if m2 != nil && maxNode != nil {
		//link minNode with maxNode
		m2.Next = maxNode
	}

	return minNode
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

func getValFromListNode(head *ListNode) []int {
	vals := make([]int, 0)
	p := head
	for p != nil {
		vals = append(vals, p.Val)
		p = p.Next
	}
	return vals
}

func main() {
	head := createListNode([]int{1, 4, 3, 2, 5, 2})
	//head := createListNode([]int{1}) //1
	head = partition(head, 3)

	vals := getValFromListNode(head)
	fmt.Println(vals)
}
