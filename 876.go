package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	return slow
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
	p := middleNode(createListNode([]int{1, 2, 3, 4, 5}))
	fmt.Println(p.Val)
	p = middleNode(createListNode([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(p.Val)
	p = middleNode(createListNode([]int{1, 2, 3}))
	fmt.Println(p.Val)
	p = middleNode(createListNode([]int{1, 2}))
	fmt.Println(p.Val)
	p = middleNode(createListNode([]int{1}))
	fmt.Println(p.Val)
	p = middleNode(createListNode([]int{}))
	fmt.Println(p)
}
